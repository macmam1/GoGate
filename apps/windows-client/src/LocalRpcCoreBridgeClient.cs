using System.Collections.Concurrent;
using System.Text;
using System.Text.Json;

namespace GoGate.WindowsClient;

public class LocalRpcCoreBridgeClient : ICoreBridgeClient, IDisposable
{
    private readonly HttpClient _http;
    private readonly CoreBridgeTransportOptions _opt;
    private readonly ConcurrentDictionary<Guid, Action<BridgeEvent>> _subs = new();
    private CancellationTokenSource? _pollCts;
    private Task? _pollTask;

    public LocalRpcCoreBridgeClient(CoreBridgeTransportOptions? options = null)
    {
        _opt = options ?? new CoreBridgeTransportOptions();
        _http = new HttpClient
        {
            BaseAddress = new Uri(_opt.BaseUrl.TrimEnd('/') + "/"),
            Timeout = TimeSpan.FromMilliseconds(_opt.RequestTimeoutMs)
        };
    }

    public ConnectResponse Connect(ConnectRequest request)
    {
        var res = SendCommand<ConnectResponse>("connect", new { profile_id = request.ProfileId, mode = request.Mode });
        return res ?? new ConnectResponse(string.Empty, "failed");
    }

    public DisconnectResponse Disconnect(DisconnectRequest request)
    {
        var res = SendCommand<DisconnectResponse>("disconnect", new { session_id = request.SessionId });
        return res ?? new DisconnectResponse("failed");
    }

    public RankCandidatesResponse RankCandidates(string mode)
    {
        var res = SendCommand<RankCandidatesResponse>("rank_candidates", new { mode });
        return res ?? new RankCandidatesResponse(Array.Empty<RankedCandidate>());
    }

    public FetchLogsResponse FetchLogs(string sessionId, string level)
    {
        var res = SendCommand<FetchLogsResponse>("fetch_logs", new { session_id = sessionId, level });
        return res ?? new FetchLogsResponse(Array.Empty<LogEvent>());
    }

    public IDisposable Subscribe(Action<BridgeEvent> onEvent)
    {
        var id = Guid.NewGuid();
        _subs[id] = onEvent;
        EnsurePolling();
        return new Subscription(() =>
        {
            _subs.TryRemove(id, out _);
            if (_subs.IsEmpty)
            {
                StopPolling();
            }
        });
    }

    private T? SendCommand<T>(string command, object payload)
    {
        var body = JsonSerializer.Serialize(new RpcCommandRequest(command, payload));
        var req = new StringContent(body, Encoding.UTF8, "application/json");
        using var resp = _http.PostAsync("command", req).GetAwaiter().GetResult();
        var raw = resp.Content.ReadAsStringAsync().GetAwaiter().GetResult();
        if (!resp.IsSuccessStatusCode)
            return default;

        var parsed = JsonSerializer.Deserialize<RpcCommandResponse<T>>(raw, JsonOpts());
        if (parsed is null || !parsed.Ok)
            return default;
        return parsed.Data;
    }

    private void EnsurePolling()
    {
        if (_pollTask is { IsCompleted: false })
            return;

        _pollCts = new CancellationTokenSource();
        _pollTask = Task.Run(() => PollLoop(_pollCts.Token));
    }

    private async Task PollLoop(CancellationToken ct)
    {
        while (!ct.IsCancellationRequested)
        {
            try
            {
                using var resp = await _http.GetAsync("events/poll", ct);
                if (resp.IsSuccessStatusCode)
                {
                    var raw = await resp.Content.ReadAsStringAsync(ct);
                    var parsed = JsonSerializer.Deserialize<EventPollResponse>(raw, JsonOpts());
                    if (parsed is { Ok: true, Events: not null })
                    {
                        foreach (var ev in parsed.Events)
                        {
                            foreach (var sub in _subs.Values)
                                sub(ev);
                        }
                    }
                }
            }
            catch
            {
                // keep polling; bridge may be temporarily unavailable
            }

            try
            {
                await Task.Delay(_opt.PollIntervalMs, ct);
            }
            catch
            {
                return;
            }
        }
    }

    private void StopPolling()
    {
        if (_pollCts == null)
            return;
        _pollCts.Cancel();
        _pollCts.Dispose();
        _pollCts = null;
        _pollTask = null;
    }

    private static JsonSerializerOptions JsonOpts() => new()
    {
        PropertyNameCaseInsensitive = true
    };

    public void Dispose()
    {
        StopPolling();
        _http.Dispose();
    }

    private record RpcCommandRequest(string Command, object Payload);

    private record RpcCommandResponse<T>(bool Ok, T? Data, string? Error);

    private record EventPollResponse(bool Ok, List<BridgeEvent>? Events, string? Error);

    private class Subscription : IDisposable
    {
        private readonly Action _onDispose;
        private bool _disposed;

        public Subscription(Action onDispose) => _onDispose = onDispose;

        public void Dispose()
        {
            if (_disposed) return;
            _disposed = true;
            _onDispose();
        }
    }
}
