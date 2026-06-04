using System;

namespace GoGate.WindowsClient;

public class MainWindowShell
{
    public string CurrentState { get; private set; } = "idle";

    public event Action<string>? OnStateChanged;

    public void Connect(string profileId, string mode)
    {
        // TODO: bind to local core bridge API
        CurrentState = "connecting";
        OnStateChanged?.Invoke(CurrentState);
    }

    public void MarkConnected()
    {
        CurrentState = "connected";
        OnStateChanged?.Invoke(CurrentState);
    }

    public void Disconnect()
    {
        // TODO: bind to local core bridge API
        CurrentState = "idle";
        OnStateChanged?.Invoke(CurrentState);
    }
}
