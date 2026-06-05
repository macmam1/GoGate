package com.gogate.android

import java.time.Instant
import java.util.UUID
import java.util.concurrent.ConcurrentHashMap

class MockCoreBridgeClient : CoreBridgeClient {
    private val listeners = ConcurrentHashMap<String, (BridgeEvent) -> Unit>()

    override fun connect(request: ConnectRequest): ConnectResponse {
        val sessionId = "sess-${UUID.randomUUID().toString().replace("-", "")}" 
        emit(BridgeEvent("session_state_changed", mapOf("state" to "connecting", "profile_id" to request.profileId)))
        emit(BridgeEvent("session_state_changed", mapOf("state" to "connected", "session_id" to sessionId, "mode" to request.mode)))
        return ConnectResponse(sessionId, "connected")
    }

    override fun disconnect(request: DisconnectRequest): DisconnectResponse {
        emit(BridgeEvent("session_state_changed", mapOf("state" to "idle", "session_id" to request.sessionId)))
        return DisconnectResponse("idle")
    }

    override fun rankCandidates(mode: String): RankCandidatesResponse {
        val list = listOf(
            RankedCandidate("profile-a", 0.89, "fast handshake path"),
            RankedCandidate("profile-b", 0.72, "balanced score")
        )
        emit(BridgeEvent("ranking_updated", mapOf("count" to list.size.toString(), "mode" to mode)))
        return RankCandidatesResponse(list)
    }

    override fun fetchLogs(sessionId: String, level: String): FetchLogsResponse {
        return FetchLogsResponse(listOf(LogEvent(Instant.now().toString(), level, "Session $sessionId log sample")))
    }

    override fun subscribe(listener: (BridgeEvent) -> Unit): DisposableSubscription {
        val id = UUID.randomUUID().toString()
        listeners[id] = listener
        return DisposableSubscription { listeners.remove(id) }
    }

    private fun emit(event: BridgeEvent) {
        listeners.values.forEach { it(event) }
    }
}
