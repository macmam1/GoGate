package engineadapters

import (
	"fmt"
	"strings"
	"time"

	"github.com/macmam1/GoGate/core/internal/contracts"
)

type XrayAdapter struct {
	binaryPath string
	sup        *Supervisor
	lastErr    error
}

func NewXrayAdapter(binaryPath string, baseArgs []string) *XrayAdapter {
	spec := ProcessSpec{Name: binaryPath, Args: append([]string{}, baseArgs...)}
	return &XrayAdapter{binaryPath: binaryPath, sup: NewSupervisor(spec)}
}

func (a *XrayAdapter) Capabilities() AdapterCapabilities {
	return AdapterCapabilities{EngineID: "xray", Protocols: []string{"vmess", "vless", "trojan", "shadowsocks"}, SupportsRouteOverride: true}
}

func (a *XrayAdapter) Start(req contracts.SessionStartRequest) (SessionHandle, error) {
	args, err := buildXrayArgs(req)
	if err != nil {
		a.lastErr = MapEngineError("xray", err, "")
		return SessionHandle{}, a.lastErr
	}

	a.sup.UpdateSpec(ProcessSpec{
		Name:    a.binaryPath,
		Args:    args,
		WorkDir: req.WorkingDir,
		Env:     mapEnv(req.Env),
	})

	if err := a.sup.Start(); err != nil {
		a.lastErr = MapEngineError("xray", err, a.sup.LastStderr())
		return SessionHandle{}, a.lastErr
	}
	return SessionHandle{ID: fmt.Sprintf("xray-%d", time.Now().UnixNano()), EngineID: "xray"}, nil
}

func (a *XrayAdapter) Health(handle SessionHandle) (HealthSnapshot, error) {
	if a.sup.Running() {
		return HealthSnapshot{Status: "running", LatencyMs: 0, HandshakeOK: true}, nil
	}
	if exitErr := a.sup.LastExitError(); exitErr != nil {
		return HealthSnapshot{Status: "crashed", LatencyMs: 0, HandshakeOK: false}, MapEngineError("xray", exitErr, a.sup.LastStderr())
	}
	return HealthSnapshot{Status: a.sup.LastStatus(), LatencyMs: 0, HandshakeOK: false}, nil
}

func (a *XrayAdapter) Stop(handle SessionHandle) error {
	err := a.sup.Stop()
	if err != nil {
		a.lastErr = MapEngineError("xray", err, a.sup.LastStderr())
		return a.lastErr
	}
	return nil
}

func (a *XrayAdapter) LastError() error { return a.lastErr }

func buildXrayArgs(req contracts.SessionStartRequest) ([]string, error) {
	cfg := strings.TrimSpace(req.RuntimeConfigPath)
	if cfg == "" {
		return nil, fmt.Errorf("missing runtime config path")
	}
	args := []string{"run", "-c", cfg}
	if req.ConnectTimeoutMs > 0 {
		args = append(args, "--timeout", fmt.Sprintf("%d", req.ConnectTimeoutMs))
	}
	if req.AdvancedModeEnable && req.Hints.HostSNI != "" {
		args = append(args, "--sni", req.Hints.HostSNI)
	}
	args = append(args, req.AdditionalArgs...)
	return args, nil
}
