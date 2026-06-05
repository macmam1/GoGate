namespace GoGate.WindowsClient;

public interface ICoreBridgeClient
{
    ConnectResponse Connect(ConnectRequest request);
    DisconnectResponse Disconnect(DisconnectRequest request);
    RankCandidatesResponse RankCandidates(string mode);
    FetchLogsResponse FetchLogs(string sessionId, string level);

    IDisposable Subscribe(Action<BridgeEvent> onEvent);
}
