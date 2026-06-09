using System;
using System.Collections.Generic;
using System.Linq;

namespace GoGate.WindowsClient;

public class MainWindowShell : IDisposable
{
    private readonly ICoreBridgeClient _bridge;
    private readonly ThemeRuntime _themeRuntime;
    private readonly NavigationHost _navigationHost;
    private IDisposable? _subscription;

    private List<ProfileItem> _profiles = new();

    public string CurrentState { get; private set; } = "idle";
    public string CurrentSessionId { get; private set; } = string.Empty;
    public string SelectedProfileId { get; private set; } = string.Empty;
    public ThemeRuntimeSettings ThemeSettings { get; private set; } = new("standard", 16, 1200, false);
    public string CurrentRoute => _navigationHost.CurrentRoute;

    public bool BridgeHealthy { get; private set; } = true;
    public string LastBridgeReason { get; private set; } = "ok";

    public event Action<string>? OnStateChanged;
    public event Action<string>? OnRouteChanged;
    public event Action<bool, string>? OnBridgeHealthChanged;

    public MainWindowShell(ICoreBridgeClient? bridge = null, ThemeRuntime? themeRuntime = null, NavigationHost? navigationHost = null)
    {
        _bridge = bridge ?? new LocalRpcCoreBridgeClient();
        _themeRuntime = themeRuntime ?? new ThemeRuntime();
        _navigationHost = navigationHost ?? new NavigationHost();

        ThemeSettings = _themeRuntime.CurrentSettings;
        _navigationHost.OnRouteChanged += route => OnRouteChanged?.Invoke(route);
        _subscription = _bridge.Subscribe(HandleBridgeEvent);
    }

    public IReadOnlyList<ProfileItem> Profiles() => _profiles;

    public void SetProfiles(IEnumerable<ProfileItem> profiles)
    {
        _profiles = profiles?.ToList() ?? new List<ProfileItem>();
        if (string.IsNullOrWhiteSpace(SelectedProfileId) && _profiles.Count > 0)
            SelectedProfileId = _profiles[0].Id;
    }

    public bool SelectProfile(string profileId)
    {
        if (_profiles.Any(p => p.Id == profileId))
        {
            SelectedProfileId = profileId;
            return true;
        }
        return false;
    }

    public void QuickConnect(string mode)
    {
        if (string.IsNullOrWhiteSpace(SelectedProfileId))
            return;
        Connect(SelectedProfileId, mode);
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

    public SessionDetails GetSessionDetails() => new(CurrentSessionId, CurrentState, BridgeHealthy, LastBridgeReason);

    public FetchLogsResponse GetLogs(string level = "info")
    {
        if (string.IsNullOrWhiteSpace(CurrentSessionId))
            return new FetchLogsResponse(Array.Empty<LogEvent>());
        return _bridge.FetchLogs(CurrentSessionId, level);
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

    public bool Navigate(string route) => _navigationHost.Navigate(route);

    private void HandleBridgeEvent(BridgeEvent ev)
    {
        if (ev.Name == "session_state_changed" && ev.Fields.TryGetValue("state", out var state))
        {
            CurrentState = state;
            OnStateChanged?.Invoke(CurrentState);
            return;
        }

        if (ev.Name == "bridge_health_changed")
        {
            var healthy = ev.Fields.TryGetValue("healthy", out var raw) && raw == "true";
            var reason = ev.Fields.TryGetValue("reason", out var r) ? r : "unknown";
            BridgeHealthy = healthy;
            LastBridgeReason = reason;
            OnBridgeHealthChanged?.Invoke(healthy, reason);
        }
    }

    public void Dispose()
    {
        _subscription?.Dispose();
        _subscription = null;
        if (_bridge is IDisposable d) d.Dispose();
    }
}
