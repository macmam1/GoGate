namespace GoGate.WindowsClient;

public enum ThemeProfile
{
    Lite,
    Balanced,
    Rich
}

public record ThemeRuntimeSettings(string Animations, int FrameBudgetMs, int BackgroundPollMs, bool ReducedMotion);

public class ThemeRuntime
{
    public ThemeProfile CurrentProfile { get; private set; } = ThemeProfile.Balanced;
    public ThemeRuntimeSettings CurrentSettings { get; private set; } =
        new("standard", 16, 1200, false);

    public ThemeRuntimeSettings Apply(ThemeProfile profile, bool reducedMotion = false)
    {
        CurrentProfile = profile;
        CurrentSettings = profile switch
        {
            ThemeProfile.Lite => new ThemeRuntimeSettings("minimal", 16, 2000, reducedMotion),
            ThemeProfile.Rich => new ThemeRuntimeSettings("advanced", 16, 800, reducedMotion),
            _ => new ThemeRuntimeSettings("standard", 16, 1200, reducedMotion),
        };
        return CurrentSettings;
    }
}
