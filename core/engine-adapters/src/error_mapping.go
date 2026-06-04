package engineadapters

import (
    "errors"
    "fmt"
    "os"
    "strings"
)

type ErrorKind string

const (
    ErrUnknown        ErrorKind = "unknown"
    ErrBinaryMissing  ErrorKind = "binary-missing"
    ErrPermission     ErrorKind = "permission"
    ErrAlreadyRunning ErrorKind = "already-running"
    ErrConfigInvalid  ErrorKind = "config-invalid"
    ErrPortInUse      ErrorKind = "port-in-use"
    ErrNetworkDenied  ErrorKind = "network-denied"
)

type AdapterError struct {
    Engine  string
    Kind    ErrorKind
    Message string
    Cause   error
}

func (e AdapterError) Error() string {
    if e.Cause == nil {
        return fmt.Sprintf("%s[%s]: %s", e.Engine, e.Kind, e.Message)
    }
    return fmt.Sprintf("%s[%s]: %s (%v)", e.Engine, e.Kind, e.Message, e.Cause)
}

func (e AdapterError) Unwrap() error { return e.Cause }

func MapEngineError(engine string, err error, stderr string) error {
    if err == nil {
        return nil
    }
    kind := classifyError(err, stderr)
    return AdapterError{Engine: engine, Kind: kind, Message: shortMessage(kind), Cause: err}
}

func classifyError(err error, stderr string) ErrorKind {
    if errors.Is(err, os.ErrNotExist) {
        return ErrBinaryMissing
    }
    if kind := semanticKindFromStderr(stderr); kind != ErrUnknown {
        return kind
    }
    lower := strings.ToLower(err.Error() + " " + stderr)
    switch {
    case strings.Contains(lower, "already running"):
        return ErrAlreadyRunning
    case strings.Contains(lower, "permission denied"):
        return ErrPermission
    case strings.Contains(lower, "invalid config"), strings.Contains(lower, "failed to parse"), strings.Contains(lower, "syntax error"):
        return ErrConfigInvalid
    case strings.Contains(lower, "address already in use"), strings.Contains(lower, "bind"):
        return ErrPortInUse
    case strings.Contains(lower, "network unreachable"), strings.Contains(lower, "no route to host"), strings.Contains(lower, "forbidden"):
        return ErrNetworkDenied
    default:
        return ErrUnknown
    }
}

func shortMessage(kind ErrorKind) string {
    switch kind {
    case ErrBinaryMissing:
        return "engine binary not found"
    case ErrPermission:
        return "permission denied while starting engine"
    case ErrAlreadyRunning:
        return "engine process already running"
    case ErrConfigInvalid:
        return "runtime config is invalid"
    case ErrPortInUse:
        return "required local port is in use"
    case ErrNetworkDenied:
        return "network operation denied or unreachable"
    default:
        return "engine runtime error"
    }
}
