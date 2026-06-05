package com.gogate.android

class MainActivityShell(
    private val bridge: CoreBridgeClient = LocalRpcCoreBridgeClient(),
    private val themeRuntime: ThemeRuntime = ThemeRuntime()
) {
    var currentState: String = "idle"
        private set

    var currentSessionId: String = ""
        private set

    var themeSettings: ThemeRuntimeSettings = themeRuntime.currentSettings
        private set

    private val subscription = bridge.subscribe { event ->
        if (event.name == "session_state_changed") {
            event.fields["state"]?.let { currentState = it }
        }
    }

    fun connect(profileId: String, mode: String) {
        currentState = "connecting"
        val response = bridge.connect(ConnectRequest(profileId, mode))
        currentSessionId = response.sessionId
        currentState = response.state
    }

    fun disconnect() {
        if (currentSessionId.isBlank()) {
            currentState = "idle"
            return
        }
        val response = bridge.disconnect(DisconnectRequest(currentSessionId))
        currentState = response.state
        currentSessionId = ""
    }

    fun applyThemeProfile(profile: String, reducedMotion: Boolean = false): ThemeRuntimeSettings {
        val parsed = when (profile.trim().lowercase()) {
            "lite" -> ThemeProfile.LITE
            "rich" -> ThemeProfile.RICH
            else -> ThemeProfile.BALANCED
        }
        themeSettings = themeRuntime.apply(parsed, reducedMotion)
        return themeSettings
    }

    fun close() {
        subscription.dispose()
    }
}
