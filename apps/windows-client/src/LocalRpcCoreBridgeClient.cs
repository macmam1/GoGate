using System.Collections.Concurrent;
using System.Text;
using System.Text.Json;
using System.Linq;

namespace GoGate.WindowsClient;

public class LocalRpcCoreBridgeClient : ICoreBridgeClient, IDisposable
{
    private readonly HttpClient _http;
    private readonly CoreBridgeTransportOptions _opt;
    private readonly ConcurrentDictionary<Guid, Action<BridgeEvent>> _subs = new();
    private CancellationTokenSource? _pollCts;
    private Task? _pollTask;

    private int _pollFailureCount;
    private int _currentBackoffMs;
    private bool _bridgeHealthy = true;
    private string _lastBridgeReason = "ok";
    private string _lastEventFingerprint = string.Empty;

    public LocalRpcCoreBridgeClient(CoreBridgeTransportOptions? options = null)
    {
        _opt = options ?? new CoreBridgeTransportOptions();
        _currentBackoffMs = _opt.PollIntervalMs;
        _http = new HttpClient
        {
            BaseAddress = new Uri(_opt.BaseUrl.TrimEnd('/') + "/"),
            Timeout = TimeSpan.FromMilliseconds(_opt.RequestTimeoutMs)
        };
    }

    public ConnectResponse Connect(ConnectRequest request)
    {
        var res = SendCommand<ConnectResponse>("connect", new { profile_id = request.ProfileId, mode = request.Mode }, "connect");
        return res ?? new ConnectResponse(string.Empty, "failed");
    }

    public DisconnectResponse Disconnect(DisconnectRequest request)
    {
        var res = SendCommand<DisconnectResponse>("disconnect", new { session_id = request.SessionId }, "disconnect");
        return res ?? new DisconnectResponse("failed");
    }

    public RankCandidatesResponse RankCandidates(string mode)
    {
        var res = SendCommand<RankCandidatesResponse>("rank_candidates", new { mode }, "rank_candidates");
        return res ?? new RankCandidatesResponse(Array.Empty<RankedCandidate>());
    }

    public FetchLogsResponse FetchLogs(string sessionId, string level)
    {
        var res = SendCommand<FetchLogsResponse>("fetch_logs", new { session_id = sessionId, level }, "fetch_logs");
        return res ?? new FetchLogsResponse(Array.Empty<LogEvent>());
    }

    public IDisposable Subscribe(Action<BridgeEvent> onEvent)
    {
        var id = Guid.NewGuid();
        _subs[id] = onEvent;
        EnsurePolling();

        // immediate health snapshot for new subscriber
        onEvent(new BridgeEvent("bridge_health_changed", new Dictionary<string, string>
        {
            ["healthy"] = _bridgeHealthy ? "true" : "false",
            ["reason"] = _lastBridgeReason
        }));

        return new Subscription(() =>
        {
            _subs.TryRemove(id, out _);
            if (_subs.IsEmpty)
            {
                StopPolling();
            }
        });
    }

    private T? SendCommand<T>(string command, object payload, string opName)
    {
        try
        {
            var body = JsonSerializer.Serialize(new RpcCommandRequest(command, payload));
            using var req = new StringContent(body, Encoding.UTF8, "application/json");
            using var resp = _http.PostAsync("command", req).GetAwaiter().GetResult();
            var raw = resp.Content.ReadAsStringAsync().GetAwaiter().GetResult();

            if (!resp.IsSuccessStatusCode)
            {
                MarkBridgeUnhealthy($"{opName}: http-{(int)resp.StatusCode}");
                return default;
            }

            var parsed = JsonSerializer.Deserialize<RpcCommandResponse<T>>(raw, JsonOpts());
            if (parsed is null || !parsed.Ok)
            {
                MarkBridgeUnhealthy($"{opName}: bad-response");
                return default;
            }

            MarkBridgeHealthy("ok");
            return parsed.Data;
        }
        catch (Exception ex)
        {
            MarkBridgeUnhealthy($"{opName}: {ex.GetType().Name}");
            return default;
        }
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
                        _pollFailureCount = 0;
                        _currentBackoffMs = _opt.PollIntervalMs;
                        MarkBridgeHealthy("ok");

                        foreach (var ev in parsed.Events)
                            BroadcastEventDedup(ev);
                    }
                    else
                    {
                        RegisterPollFailure("poll: invalid payload");
                    }
                }
                else
                {
                    RegisterPollFailure($"poll: http-{(int)resp.StatusCode}");
                }
            }
            catch (Exception ex)
            {
                RegisterPollFailure($"poll: {ex.GetType().Name}");
            }

            var delay = _pollFailureCount > 0 ? _currentBackoffMs : _opt.PollIntervalMs;
            try
            {
                await Task.Delay(delay, ct);
            }
            catch
            {
                return;
            }
        }
    }

    private void RegisterPollFailure(string reason)
    {
        _pollFailureCount++;
        if (_pollFailureCount >= _opt.FailureThresholdForDegraded)
            MarkBridgeUnhealthy(reason);

        var next = _currentBackoffMs <= 0 ? _opt.PollIntervalMs : _currentBackoffMs * 2;
        _currentBackoffMs = Math.Min(next, _opt.MaxBackoffMs);
    }

    private void MarkBridgeHealthy(string reason)
    {
        if (_bridgeHealthy && _lastBridgeReason == reason)
            return;

        _bridgeHealthy = true;
        _lastBridgeReason = reason;
        BroadcastEvent(new BridgeEvent("bridge_health_changed", new Dictionary<string, string>
        {
            ["healthy"] = "true",
            ["reason"] = reason
        }));
    }

    private void MarkBridgeUnhealthy(string reason)
    {
        if (!_bridgeHealthy && _lastBridgeReason == reason)
            return;

        _bridgeHealthy = false;
        _lastBridgeReason = reason;
        BroadcastEvent(new BridgeEvent("bridge_health_changed", new Dictionary<string, string>
        {
            ["healthy"] = "false",
            ["reason"] = reason
        }));
    }

    private void BroadcastEvent(BridgeEvent ev)
    {
        foreach (var sub in _subs.Values)
            sub(ev);
    }

    private void BroadcastEventDedup(BridgeEvent ev)
    {
        var fingerprint = $"{ev.Name}:{string.Join(';', ev.Fields.OrderBy(x => x.Key).Select(x => x.Key + "=" + x.Value))}";
        if (fingerprint == _lastEventFingerprint)
            return;
        _lastEventFingerprint = fingerprint;
        BroadcastEvent(ev);
    }

    private void StopPolling()
    {
        if (_pollCts == null)
            return;
        _pollCts.Cancel();
        _pollCts.Dispose();
        _pollCts = null;
        _pollTask = null;
        _pollFailureCount = 0;
        _currentBackoffMs = _opt.PollIntervalMs;
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
