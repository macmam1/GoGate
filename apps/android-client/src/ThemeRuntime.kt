package com.gogate.android

enum class ThemeProfile { LITE, BALANCED, RICH }

data class ThemeRuntimeSettings(
    val animations: String,
    val frameBudgetMs: Int,
    val backgroundPollMs: Int,
    val reducedMotion: Boolean
)

class ThemeRuntime {
    var currentProfile: ThemeProfile = ThemeProfile.BALANCED
        private set

    var currentSettings: ThemeRuntimeSettings =
        ThemeRuntimeSettings("standard", 16, 1200, false)
        private set

    fun apply(profile: ThemeProfile, reducedMotion: Boolean = false): ThemeRuntimeSettings {
        currentProfile = profile
        currentSettings = when (profile) {
            ThemeProfile.LITE -> ThemeRuntimeSettings("minimal", 16, 2000, reducedMotion)
            ThemeProfile.RICH -> ThemeRuntimeSettings("advanced", 16, 800, reducedMotion)
            ThemeProfile.BALANCED -> ThemeRuntimeSettings("standard", 16, 1200, reducedMotion)
        }
        return currentSettings
    }
}
