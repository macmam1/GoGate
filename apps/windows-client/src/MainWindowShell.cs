using System;

namespace GoGate.WindowsClient;

public class MainWindowShell : IDisposable
{
    private readonly ICoreBridgeClient _bridge;
    private IDisposable? _subscription;

    public string CurrentState { get; private set; } = "idle";
    public string CurrentSessionId { get; private set; } = string.Empty;

    public event Action<string>? OnStateChanged;

    public MainWindowShell(ICoreBridgeClient? bridge = null)
    {
        _bridge = bridge ?? new MockCoreBridgeClient();
        _subscription = _bridge.Subscribe(HandleBridgeEvent);
    }

    public void Connect(string profileId, string mode)
    {
        CurrentState = "connecting";
        OnStateChanged?.Invoke(CurrentState);

        var response = _bridge.Connect(new ConnectRequest(profileId, mode));
        CurrentSessionId = response.SessionId;
        CurrentState = response.State;
        OnStateChanged?.Invoke(CurrentState);
    }

    public void Disconnect()
    {
        if (string.IsNullOrWhiteSpace(CurrentSessionId))
        {
            CurrentState = "idle";
            OnStateChanged?.Invoke(CurrentState);
            return;
        }

        var response = _bridge.Disconnect(new DisconnectRequest(CurrentSessionId));
        CurrentState = response.State;
        CurrentSessionId = string.Empty;
        OnStateChanged?.Invoke(CurrentState);
    }

    private void HandleBridgeEvent(BridgeEvent ev)
    {
        if (ev.Name == "session_state_changed" && ev.Fields.TryGetValue("state", out var state))
        {
            CurrentState = state;
            OnStateChanged?.Invoke(CurrentState);
        }
    }

    public void Dispose()
    {
        _subscription?.Dispose();
        _subscription = null;
    }
}
