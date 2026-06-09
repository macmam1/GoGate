using System;
using System.Collections.Generic;

namespace GoGate.WindowsClient;

public class NavigationHost
{
    private static readonly HashSet<string> AllowedRoutes = new(StringComparer.OrdinalIgnoreCase)
    {
        "home", "settings", "ip_scanner", "logs", "about"
    };

    public string CurrentRoute { get; private set; } = "home";

    public event Action<string>? OnRouteChanged;

    public bool Navigate(string route)
    {
        var normalized = route?.Trim().ToLowerInvariant() ?? string.Empty;
        if (!AllowedRoutes.Contains(normalized))
            return false;
        if (CurrentRoute == normalized)
            return true;

        CurrentRoute = normalized;
        OnRouteChanged?.Invoke(CurrentRoute);
        return true;
    }
}
