package engineadapters

import (
	"errors"
	"os/exec"
	"testing"
	"time"
)

func TestSupervisorCrashFixtureMapping(t *testing.T) {
	if _, err := exec.LookPath("sh"); err != nil {
		t.Skip("sh not available")
	}
	s := NewSupervisor(ProcessSpec{Name: "sh", Args: []string{"-c", "echo 'address already in use' 1>&2; exit 1"}})
	if err := s.Start(); err != nil {
		t.Fatalf("start fixture: %v", err)
	}

	deadline := time.Now().Add(2 * time.Second)
	for s.Running() && time.Now().Before(deadline) {
		time.Sleep(20 * time.Millisecond)
	}

	mapped := MapEngineError("xray", s.LastExitError(), s.LastStderr())
	var ae AdapterError
	if !errors.As(mapped, &ae) {
		t.Fatalf("expected adapter error")
	}
	if ae.Kind != ErrPortInUse {
		t.Fatalf("expected ErrPortInUse, got %s", ae.Kind)
	}
}
