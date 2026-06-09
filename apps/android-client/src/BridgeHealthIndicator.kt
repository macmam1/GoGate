package com.gogate.android

data class BridgeHealthViewModel(
    val healthy: Boolean,
    val reason: String,
    val badgeColor: String,
    val label: String
)

object BridgeHealthIndicator {
    fun from(healthy: Boolean, reason: String): BridgeHealthViewModel {
        return if (healthy) {
            BridgeHealthViewModel(true, reason, "#2ECC71", "Bridge Healthy")
        } else {
            val lowered = reason.lowercase()
            val color = if (lowered.contains("timeout") || lowered.contains("poll")) "#F39C12" else "#E74C3C"
            BridgeHealthViewModel(false, reason, color, "Bridge Degraded")
        }
    }
}
