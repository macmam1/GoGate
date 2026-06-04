package probe

import (
	"context"
	"sync"
)

type WorkerPool struct {
	prober      Prober
	concurrency int
}

func NewWorkerPool(prober Prober, concurrency int) *WorkerPool {
	if concurrency <= 0 {
		concurrency = 2
	}
	return &WorkerPool{prober: prober, concurrency: concurrency}
}

func (w *WorkerPool) Run(ctx context.Context, requests []Request) []Result {
	if len(requests) == 0 || w.prober == nil {
		return nil
	}

	jobs := make(chan Request)
	results := make(chan Result, len(requests))
	wg := sync.WaitGroup{}

	for i := 0; i < w.concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case req, ok := <-jobs:
					if !ok {
						return
					}
					results <- w.prober.Probe(ctx, req)
				}
			}
		}()
	}

	for _, req := range requests {
		jobs <- req
	}
	close(jobs)
	wg.Wait()
	close(results)

	out := make([]Result, 0, len(requests))
	for r := range results {
		out = append(out, r)
	}
	return out
}
