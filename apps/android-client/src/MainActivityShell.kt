package com.gogate.android

class MainActivityShell {
    var currentState: String = "idle"
        private set

    fun connect(profileId: String, mode: String) {
        // TODO: bind to core bridge command: connect
        currentState = "connecting"
    }

    fun markConnected() {
        currentState = "connected"
    }

    fun disconnect() {
        // TODO: bind to core bridge command: disconnect
        currentState = "idle"
    }
}
