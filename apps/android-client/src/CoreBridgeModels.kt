package com.gogate.android

data class ConnectRequest(val profileId: String, val mode: String)
data class ConnectResponse(val sessionId: String, val state: String)

data class DisconnectRequest(val sessionId: String)
data class DisconnectResponse(val state: String)

data class RankedCandidate(val profileId: String, val score: Double, val reason: String)
data class RankCandidatesResponse(val candidates: List<RankedCandidate>)

data class LogEvent(val ts: String, val level: String, val message: String)
data class FetchLogsResponse(val events: List<LogEvent>)

data class BridgeEvent(val name: String, val fields: Map<String, String>)
