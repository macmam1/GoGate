package harvester

import (
    "context"
    "sync/atomic"
    "testing"
    "time"
)

func TestIngestionScheduler(t *testing.T) {
    var runs int32
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    s := NewIngestionScheduler(20*time.Millisecond, func(context.Context) error {
        atomic.AddInt32(&runs, 1)
        return nil
    })
    s.Start(ctx, true)
    time.Sleep(65 * time.Millisecond)
    s.Stop()

    if atomic.LoadInt32(&runs) < 2 {
        t.Fatalf("expected scheduler to run at least twice")
    }
}

func TestIngestionSchedulerHooks(t *testing.T) {
    var before int32
    var after int32
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    s := NewIngestionSchedulerWithHooks(25*time.Millisecond, func(context.Context) error {
        return nil
    }, SchedulerHooks{
        BeforeRun: func(time.Time) { atomic.AddInt32(&before, 1) },
        AfterRun:  func(time.Time, error) { atomic.AddInt32(&after, 1) },
    })
    s.Start(ctx, true)
    time.Sleep(55 * time.Millisecond)
    s.Stop()

    if atomic.LoadInt32(&before) == 0 || atomic.LoadInt32(&after) == 0 {
        t.Fatalf("expected hooks to be called")
    }
}
