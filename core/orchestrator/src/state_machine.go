package orchestrator

type SessionState string

const (
	StateIdle       SessionState = "idle"
	StateConnecting SessionState = "connecting"
	StateConnected  SessionState = "connected"
	StateFallback   SessionState = "fallback"
	StateFailed     SessionState = "failed"
)

type SessionEvent string

const (
	EventConnectRequested SessionEvent = "connect-requested"
	EventConnectSucceeded SessionEvent = "connect-succeeded"
	EventConnectFailed    SessionEvent = "connect-failed"
	EventFallbackNext     SessionEvent = "fallback-next"
	EventDisconnected     SessionEvent = "disconnected"
)

func NextState(curr SessionState, ev SessionEvent) SessionState {
	switch curr {
	case StateIdle:
		if ev == EventConnectRequested {
			return StateConnecting
		}
	case StateConnecting:
		switch ev {
		case EventConnectSucceeded:
			return StateConnected
		case EventConnectFailed:
			return StateFallback
		}
	case StateFallback:
		switch ev {
		case EventFallbackNext:
			return StateConnecting
		case EventConnectFailed:
			return StateFailed
		}
	case StateConnected:
		if ev == EventDisconnected {
			return StateIdle
		}
	case StateFailed:
		if ev == EventConnectRequested {
			return StateConnecting
		}
	}
	return curr
}
