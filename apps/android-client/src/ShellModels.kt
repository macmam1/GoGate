package com.gogate.android

data class ProfileItem(val id: String, val title: String, val location: String, val protocol: String)

data class SessionDetails(
    val sessionId: String,
    val state: String,
    val bridgeHealthy: Boolean,
    val bridgeReason: String
)
