package harvester

import (
    "context"
    "sync"
    "time"
)

type IngestionTask func(context.Context) error

type SchedulerHooks struct {
    BeforeRun func(time.Time)
    AfterRun  func(time.Time, error)
}

type IngestionScheduler struct {
    interval time.Duration
    task     IngestionTask
    hooks    SchedulerHooks

    stopOnce sync.Once
    stopCh   chan struct{}
    wg       sync.WaitGroup
}

func NewIngestionScheduler(interval time.Duration, task IngestionTask) *IngestionScheduler {
    return NewIngestionSchedulerWithHooks(interval, task, SchedulerHooks{})
}

func NewIngestionSchedulerWithHooks(interval time.Duration, task IngestionTask, hooks SchedulerHooks) *IngestionScheduler {
    if interval <= 0 {
        interval = 5 * time.Minute
    }
    return &IngestionScheduler{interval: interval, task: task, hooks: hooks, stopCh: make(chan struct{})}
}

func (s *IngestionScheduler) Start(ctx context.Context, runImmediately bool) {
    s.wg.Add(1)
    go func() {
        defer s.wg.Done()
        ticker := time.NewTicker(s.interval)
        defer ticker.Stop()

        if runImmediately {
            s.runOnce(ctx)
        }

        for {
            select {
            case <-ctx.Done():
                return
            case <-s.stopCh:
                return
            case <-ticker.C:
                s.runOnce(ctx)
            }
        }
    }()
}

func (s *IngestionScheduler) Stop() {
    s.stopOnce.Do(func() {
        close(s.stopCh)
    })
    s.wg.Wait()
}

func (s *IngestionScheduler) runOnce(ctx context.Context) {
    now := time.Now()
    if s.hooks.BeforeRun != nil {
        s.hooks.BeforeRun(now)
    }
    var err error
    if s.task != nil {
        err = s.task(ctx)
    }
    if s.hooks.AfterRun != nil {
        s.hooks.AfterRun(now, err)
    }
}
