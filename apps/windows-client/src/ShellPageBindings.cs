using System.Collections.Generic;

namespace GoGate.WindowsClient;

public interface IHomePageView
{
    void RenderSession(SessionDetails session);
    void RenderBridgeHealth(BridgeHealthViewModel vm);
}

public interface ILogsPageView
{
    void RenderLogs(IReadOnlyList<LogEvent> events);
}

public class ShellPageBindings
{
    private readonly MainWindowShell _shell;
    private readonly IHomePageView _home;
    private readonly ILogsPageView _logs;

    public ShellPageBindings(MainWindowShell shell, IHomePageView home, ILogsPageView logs)
    {
        _shell = shell;
        _home = home;
        _logs = logs;

        _shell.OnStateChanged += _ => RenderHome();
        _shell.OnBridgeHealthChanged += (_, _) => RenderHome();
    }

    public void RenderHome()
    {
        var session = _shell.GetSessionDetails();
        _home.RenderSession(session);
        _home.RenderBridgeHealth(BridgeHealthIndicator.From(session.BridgeHealthy, session.BridgeReason));
    }

    public void RenderLogs(string level = "info")
    {
        var res = _shell.GetLogs(level);
        _logs.RenderLogs(res.Events);
    }
}
