package com.gogate.android

interface CoreBridgeClient {
    fun connect(request: ConnectRequest): ConnectResponse
    fun disconnect(request: DisconnectRequest): DisconnectResponse
    fun rankCandidates(mode: String): RankCandidatesResponse
    fun fetchLogs(sessionId: String, level: String): FetchLogsResponse

    fun subscribe(listener: (BridgeEvent) -> Unit): DisposableSubscription
}

class DisposableSubscription(private val onDispose: () -> Unit) {
    private var disposed = false

    fun dispose() {
        if (disposed) return
        disposed = true
        onDispose()
    }
}
