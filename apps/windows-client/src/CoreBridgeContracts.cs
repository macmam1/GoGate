namespace GoGate.WindowsClient;

public record ConnectRequest(string ProfileId, string Mode);
public record ConnectResponse(string SessionId, string State);

public record DisconnectRequest(string SessionId);
public record DisconnectResponse(string State);

public record RankedCandidate(string ProfileId, double Score, string Reason);
public record RankCandidatesResponse(IReadOnlyList<RankedCandidate> Candidates);

public record LogEvent(string Ts, string Level, string Message);
public record FetchLogsResponse(IReadOnlyList<LogEvent> Events);

public record BridgeEvent(string Name, IReadOnlyDictionary<string, string> Fields);
