package com.gogate.android

import org.json.JSONArray
import org.json.JSONObject

object SharedProfileConfig {
    fun parse(json: String): List<ProfileItem> {
        val root = JSONObject(json)
        if (root.optString("version") != "v1") return emptyList()
        val arr = root.optJSONArray("profiles") ?: JSONArray()
        val out = mutableListOf<ProfileItem>()
        for (i in 0 until arr.length()) {
            val o = arr.optJSONObject(i) ?: continue
            out += ProfileItem(
                id = o.optString("id", ""),
                title = o.optString("title", ""),
                location = o.optString("location", ""),
                protocol = o.optString("protocol", "")
            )
        }
        return out
    }

    fun serialize(profiles: List<ProfileItem>): String {
        val arr = JSONArray()
        profiles.forEach {
            arr.put(JSONObject()
                .put("id", it.id)
                .put("title", it.title)
                .put("location", it.location)
                .put("protocol", it.protocol))
        }
        return JSONObject()
            .put("version", "v1")
            .put("profiles", arr)
            .toString(2)
    }
}
