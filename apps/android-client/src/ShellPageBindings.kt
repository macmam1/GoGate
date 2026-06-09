package com.gogate.android

interface HomePageView {
    fun renderSession(session: SessionDetails)
    fun renderBridgeHealth(viewModel: BridgeHealthViewModel)
}

interface LogsPageView {
    fun renderLogs(events: List<LogEvent>)
}

class ShellPageBindings(
    private val shell: MainActivityShell,
    private val home: HomePageView,
    private val logs: LogsPageView
) {
    fun renderHome() {
        val session = shell.getSessionDetails()
        home.renderSession(session)
        home.renderBridgeHealth(BridgeHealthIndicator.from(session.bridgeHealthy, session.bridgeReason))
    }

    fun renderLogs(level: String = "info") {
        logs.renderLogs(shell.getLogs(level).events)
    }
}
