package com.gogate.android

class NavigationHost {
    private val allowed = setOf("home", "server_list", "connection_details", "logs", "settings", "about")

    var currentRoute: String = "home"
        private set

    private val listeners = mutableListOf<(String) -> Unit>()

    fun navigate(route: String): Boolean {
        val normalized = route.trim().lowercase()
        if (!allowed.contains(normalized)) return false
        if (currentRoute == normalized) return true
        currentRoute = normalized
        listeners.forEach { it(currentRoute) }
        return true
    }

    fun subscribe(listener: (String) -> Unit): DisposableSubscription {
        listeners += listener
        return DisposableSubscription { listeners.remove(listener) }
    }
}
