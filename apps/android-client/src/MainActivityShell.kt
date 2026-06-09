package com.gogate.android

class MainActivityShell(
    private val bridge: CoreBridgeClient = LocalRpcCoreBridgeClient(),
    private val themeRuntime: ThemeRuntime = ThemeRuntime(),
    private val navigationHost: NavigationHost = NavigationHost()
) {
    var currentState: String = "idle"
        private set

    var currentSessionId: String = ""
        private set

    var themeSettings: ThemeRuntimeSettings = themeRuntime.currentSettings
        private set

    var currentRoute: String = navigationHost.currentRoute
        private set

    var bridgeHealthy: Boolean = true
        private set

    var lastBridgeReason: String = "ok"
        private set

    private val navSub = navigationHost.subscribe { route ->
        currentRoute = route
    }

    private val bridgeSub = bridge.subscribe { event ->
        when (event.name) {
            "session_state_changed" -> event.fields["state"]?.let { currentState = it }
            "bridge_health_changed" -> {
                bridgeHealthy = event.fields["healthy"] == "true"
                lastBridgeReason = event.fields["reason"] ?: "unknown"
            }
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

    fun navigate(route: String): Boolean = navigationHost.navigate(route)

    fun close() {
        bridgeSub.dispose()
        navSub.dispose()
        if (bridge is AutoCloseable) {
            bridge.close()
        }
    }
}
