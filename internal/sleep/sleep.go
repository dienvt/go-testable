package sleep

import (
	"context"
	"time"
)

type Sleeper struct {
	d time.Duration
}

// SleepUntil block process until reaches config duration
func (s *Sleeper) SleepUntil(_ context.Context, d time.Duration) {
	duration := s.d - d
	if duration > 0 {
		time.Sleep(duration)
	}
}

type SleeperOptimized struct {
	d time.Duration

	// sleepFunc using for sleep execution
	sleepFunc func(time.Duration)
}

// SleepUntil block process until reaches config duration
func (s *SleeperOptimized) SleepUntil(_ context.Context, d time.Duration) {
	duration := s.d - d
	if duration > 0 {
		s.sleepFunc(duration)
	}
}
