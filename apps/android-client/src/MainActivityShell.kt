package com.gogate.android

class MainActivityShell(
    private val bridge: CoreBridgeClient = MockCoreBridgeClient()
) {
    var currentState: String = "idle"
        private set

    var currentSessionId: String = ""
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

    fun close() {
        subscription.dispose()
    }
}
