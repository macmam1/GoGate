namespace GoGate.WindowsClient;

public record ProfileItem(string Id, string Title, string Location, string Protocol);

public record SessionDetails(string SessionId, string State, bool BridgeHealthy, string BridgeReason);
