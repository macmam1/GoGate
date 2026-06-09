namespace GoGate.WindowsClient;

public class CoreBridgeTransportOptions
{
    public string BaseUrl { get; init; } = "http://127.0.0.1:47831";
    public int PollIntervalMs { get; init; } = 750;
    public int RequestTimeoutMs { get; init; } = 6000;
    public int MaxBackoffMs { get; init; } = 8000;
    public int FailureThresholdForDegraded { get; init; } = 2;
}
