package engineadapters

import "testing"

func TestSupervisorLifecycle(t *testing.T) {
	s := NewSupervisor(ProcessSpec{Name: "sleep", Args: []string{"1"}})
	if err := s.Start(); err != nil {
		t.Fatalf("start failed: %v", err)
	}
	if !s.Running() {
		t.Fatalf("expected running state")
	}
	if err := s.Stop(); err != nil {
		t.Fatalf("stop failed: %v", err)
	}
}
