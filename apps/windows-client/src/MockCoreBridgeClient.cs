using System.Collections.Concurrent;

namespace GoGate.WindowsClient;

public class MockCoreBridgeClient : ICoreBridgeClient
{
    private readonly ConcurrentDictionary<Guid, Action<BridgeEvent>> _subs = new();

    public ConnectResponse Connect(ConnectRequest request)
    {
        var sessionId = $"sess-{Guid.NewGuid():N}";
        Emit(new BridgeEvent("session_state_changed", new Dictionary<string, string>
        {
            ["state"] = "connecting",
            ["profile_id"] = request.ProfileId,
            ["mode"] = request.Mode
        }));
        Emit(new BridgeEvent("session_state_changed", new Dictionary<string, string>
        {
            ["state"] = "connected",
            ["profile_id"] = request.ProfileId,
            ["mode"] = request.Mode,
            ["session_id"] = sessionId
        }));
        return new ConnectResponse(sessionId, "connected");
    }

    public DisconnectResponse Disconnect(DisconnectRequest request)
    {
        Emit(new BridgeEvent("session_state_changed", new Dictionary<string, string>
        {
            ["state"] = "idle",
            ["session_id"] = request.SessionId
        }));
        return new DisconnectResponse("idle");
    }

    public RankCandidatesResponse RankCandidates(string mode)
    {
        var list = new List<RankedCandidate>
        {
            new("profile-a", 0.91, "fast handshake path"),
            new("profile-b", 0.74, "balanced score")
        };
        Emit(new BridgeEvent("ranking_updated", new Dictionary<string, string> { ["count"] = list.Count.ToString() }));
        return new RankCandidatesResponse(list);
    }

    public FetchLogsResponse FetchLogs(string sessionId, string level)
    {
        var now = DateTimeOffset.UtcNow.ToString("O");
        return new FetchLogsResponse(new List<LogEvent>
        {
            new(now, level, $"Session {sessionId} log sample")
        });
    }

    public IDisposable Subscribe(Action<BridgeEvent> onEvent)
    {
        var id = Guid.NewGuid();
        _subs[id] = onEvent;
        return new Subscription(() => _subs.TryRemove(id, out _));
    }

    private void Emit(BridgeEvent ev)
    {
        foreach (var sub in _subs.Values)
            sub(ev);
    }

    private class Subscription : IDisposable
    {
        private readonly Action _dispose;
        private bool _done;

        public Subscription(Action dispose) => _dispose = dispose;

        public void Dispose()
        {
            if (_done) return;
            _done = true;
            _dispose();
        }
    }
}
