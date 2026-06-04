package harvester

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"
)

func TestFileSchedulerStateStore(t *testing.T) {
	store := &FileSchedulerStateStore{Path: filepath.Join(t.TempDir(), "scheduler", "state.json")}
	now := time.Now().UTC().Truncate(time.Second)
	in := SchedulerRunState{LastRunAt: now, LastStatus: "ok"}
	if err := store.Save(context.Background(), in); err != nil {
		t.Fatalf("save: %v", err)
	}
	out, err := store.Load(context.Background())
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if out.LastStatus != "ok" || out.LastRunAt.IsZero() {
		t.Fatalf("unexpected loaded state: %+v", out)
	}
}

func TestPersistentSchedulerHooks(t *testing.T) {
	store := &FileSchedulerStateStore{Path: filepath.Join(t.TempDir(), "state.json")}
	sink := &InMemoryEventSink{}
	hooks := PersistentSchedulerHooks(store, sink)
	at := time.Now()
	hooks.BeforeRun(at)
	hooks.AfterRun(at, errors.New("boom"))

	state, err := store.Load(context.Background())
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if state.LastStatus != "error" {
		t.Fatalf("expected error status, got %s", state.LastStatus)
	}
	if len(sink.Events) < 2 {
		t.Fatalf("expected scheduler events emitted")
	}
}
