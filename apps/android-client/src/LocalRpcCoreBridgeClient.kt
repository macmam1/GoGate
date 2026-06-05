package com.gogate.android

import org.json.JSONArray
import org.json.JSONObject
import java.io.BufferedReader
import java.io.InputStreamReader
import java.io.OutputStreamWriter
import java.net.HttpURLConnection
import java.net.URL
import java.time.Instant
import java.util.UUID
import java.util.concurrent.ConcurrentHashMap
import kotlin.concurrent.thread

class LocalRpcCoreBridgeClient(
    private val baseUrl: String = "http://127.0.0.1:47831",
    private val pollIntervalMs: Long = 750L
) : CoreBridgeClient {

    private val listeners = ConcurrentHashMap<String, (BridgeEvent) -> Unit>()
    @Volatile private var polling = false

    override fun connect(request: ConnectRequest): ConnectResponse {
        val payload = JSONObject()
            .put("profile_id", request.profileId)
            .put("mode", request.mode)
        val data = sendCommand("connect", payload)
        return if (data != null) {
            ConnectResponse(data.optString("session_id", ""), data.optString("state", "failed"))
        } else {
            ConnectResponse("", "failed")
        }
    }

    override fun disconnect(request: DisconnectRequest): DisconnectResponse {
        val payload = JSONObject().put("session_id", request.sessionId)
        val data = sendCommand("disconnect", payload)
        return DisconnectResponse(data?.optString("state", "failed") ?: "failed")
    }

    override fun rankCandidates(mode: String): RankCandidatesResponse {
        val payload = JSONObject().put("mode", mode)
        val data = sendCommand("rank_candidates", payload) ?: return RankCandidatesResponse(emptyList())
        val arr = data.optJSONArray("candidates") ?: JSONArray()
        val out = mutableListOf<RankedCandidate>()
        for (i in 0 until arr.length()) {
            val it = arr.optJSONObject(i) ?: continue
            out += RankedCandidate(
                profileId = it.optString("profile_id", ""),
                score = it.optDouble("score", 0.0),
                reason = it.optString("reason", "")
            )
        }
        return RankCandidatesResponse(out)
    }

    override fun fetchLogs(sessionId: String, level: String): FetchLogsResponse {
        val payload = JSONObject().put("session_id", sessionId).put("level", level)
        val data = sendCommand("fetch_logs", payload)
        if (data == null) {
            return FetchLogsResponse(listOf(LogEvent(Instant.now().toString(), level, "log unavailable")))
        }
        val arr = data.optJSONArray("events") ?: JSONArray()
        val out = mutableListOf<LogEvent>()
        for (i in 0 until arr.length()) {
            val it = arr.optJSONObject(i) ?: continue
            out += LogEvent(
                ts = it.optString("ts", Instant.now().toString()),
                level = it.optString("level", level),
                message = it.optString("message", "")
            )
        }
        return FetchLogsResponse(out)
    }

    override fun subscribe(listener: (BridgeEvent) -> Unit): DisposableSubscription {
        val id = UUID.randomUUID().toString()
        listeners[id] = listener
        ensurePolling()
        return DisposableSubscription {
            listeners.remove(id)
            if (listeners.isEmpty()) polling = false
        }
    }

    private fun ensurePolling() {
        if (polling) return
        polling = true
        thread(start = true, isDaemon = true) {
            while (polling) {
                try {
                    val events = pollEvents()
                    events.forEach { ev -> listeners.values.forEach { it(ev) } }
                } catch (_: Exception) {
                    // bridge may be unavailable temporarily
                }
                Thread.sleep(pollIntervalMs)
            }
        }
    }

    private fun pollEvents(): List<BridgeEvent> {
        val conn = (URL("$baseUrl/events/poll").openConnection() as HttpURLConnection).apply {
            requestMethod = "GET"
            connectTimeout = 5000
            readTimeout = 5000
        }
        conn.inputStream.use { input ->
            val body = BufferedReader(InputStreamReader(input)).readText()
            val root = JSONObject(body)
            if (!root.optBoolean("ok", false)) return emptyList()
            val arr = root.optJSONArray("events") ?: JSONArray()
            val out = mutableListOf<BridgeEvent>()
            for (i in 0 until arr.length()) {
                val o = arr.optJSONObject(i) ?: continue
                val fieldsObj = o.optJSONObject("fields") ?: JSONObject()
                val fields = mutableMapOf<String, String>()
                val keys = fieldsObj.keys()
                while (keys.hasNext()) {
                    val k = keys.next()
                    fields[k] = fieldsObj.optString(k, "")
                }
                out += BridgeEvent(o.optString("name", ""), fields)
            }
            return out
        }
    }

    private fun sendCommand(command: String, payload: JSONObject): JSONObject? {
        val conn = (URL("$baseUrl/command").openConnection() as HttpURLConnection).apply {
            requestMethod = "POST"
            doOutput = true
            connectTimeout = 6000
            readTimeout = 6000
            setRequestProperty("Content-Type", "application/json")
        }
        val body = JSONObject()
            .put("command", command)
            .put("payload", payload)
            .toString()
        OutputStreamWriter(conn.outputStream).use { it.write(body) }

        val input = if (conn.responseCode < 400) conn.inputStream else conn.errorStream ?: return null
        val raw = BufferedReader(InputStreamReader(input)).readText()
        val root = JSONObject(raw)
        if (!root.optBoolean("ok", false)) return null
        return root.optJSONObject("data")
    }
}
