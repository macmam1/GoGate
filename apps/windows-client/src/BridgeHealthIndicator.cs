namespace GoGate.WindowsClient;

public record BridgeHealthViewModel(bool Healthy, string Reason, string BadgeColor, string Label);

public static class BridgeHealthIndicator
{
    public static BridgeHealthViewModel From(bool healthy, string reason)
    {
        if (healthy)
            return new BridgeHealthViewModel(true, reason, "#2ECC71", "Bridge Healthy");

        var lowered = (reason ?? string.Empty).ToLowerInvariant();
        var color = lowered.Contains("timeout") || lowered.Contains("poll") ? "#F39C12" : "#E74C3C";
        return new BridgeHealthViewModel(false, reason, color, "Bridge Degraded");
    }
}
