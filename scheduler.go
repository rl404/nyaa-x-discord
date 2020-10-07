package main

import (
	"context"
	"sync"
	"time"
)

// Job is a simple context func for scheduler job.
type Job func(ctx context.Context)

// Scheduler is model for scheduler.
type Scheduler struct {
	WG     *sync.WaitGroup
	Cancel []context.CancelFunc
}

// newScheduler to create new scheduler instance.
func newScheduler() *Scheduler {
	return &Scheduler{
		WG:     new(sync.WaitGroup),
		Cancel: make([]context.CancelFunc, 0),
	}
}

// add to add job to scheduler with interval delay..
func (s *Scheduler) add(ctx context.Context, j Job, interval time.Duration) {
	ctx, cancel := context.WithCancel(ctx)
	s.Cancel = append(s.Cancel, cancel)

	s.WG.Add(1)
	go s.process(ctx, j, interval)
}

// stop to stop all running jobs.
func (s *Scheduler) stop() {
	for _, cancel := range s.Cancel {
		cancel()
	}
	s.WG.Wait()
}

// process to run goroutine to listen to ticker.
func (s *Scheduler) process(ctx context.Context, j Job, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			j(ctx)
		case <-ctx.Done():
			s.WG.Done()
			return
		}
	}
}
