using System;

namespace GoGate.WindowsClient;

public class MainWindowShell : IDisposable
{
    private readonly ICoreBridgeClient _bridge;
    private readonly ThemeRuntime _themeRuntime;
    private IDisposable? _subscription;

    public string CurrentState { get; private set; } = "idle";
    public string CurrentSessionId { get; private set; } = string.Empty;
    public ThemeRuntimeSettings ThemeSettings { get; private set; } = new("standard", 16, 1200, false);

    public event Action<string>? OnStateChanged;

    public MainWindowShell(ICoreBridgeClient? bridge = null, ThemeRuntime? themeRuntime = null)
    {
        _bridge = bridge ?? new LocalRpcCoreBridgeClient();
        _themeRuntime = themeRuntime ?? new ThemeRuntime();
        ThemeSettings = _themeRuntime.CurrentSettings;
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

    public ThemeRuntimeSettings ApplyThemeProfile(string profile, bool reducedMotion = false)
    {
        var parsed = profile.Trim().ToLowerInvariant() switch
        {
            "lite" => ThemeProfile.Lite,
            "rich" => ThemeProfile.Rich,
            _ => ThemeProfile.Balanced,
        };
        ThemeSettings = _themeRuntime.Apply(parsed, reducedMotion);
        return ThemeSettings;
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
        if (_bridge is IDisposable d) d.Dispose();
    }
}
