package sleep

import (
	"context"
	"testing"
	"time"
)

func TestSleeper_SleepUntil(t *testing.T) {
	// log time for calculate
	start := time.Now()
	s := Sleeper{
		d: 5 * time.Second,
	}

	s.SleepUntil(context.Background(), 1*time.Second)
	executeTime := time.Since(start)
	if executeTime < 4*time.Second {
		t.Errorf("expect sleep 4 second but got %v", executeTime)
	}

	s.SleepUntil(context.Background(), 10*time.Second)
	executeTime = time.Since(start)
	if executeTime < time.Second {
		t.Errorf("expect do not sleep got %v", executeTime)
	}
}

func TestSleeperOptimized_SleepUntil(t *testing.T) {
	// don't need log time for calculate and waiting time
	var sleepDur time.Duration
	s := SleeperOptimized{
		d: 5 * time.Second,
		sleepFunc: func(duration time.Duration) {
			sleepDur = duration
		},
	}

	s.SleepUntil(context.Background(), 1*time.Second)
	if sleepDur != 4*time.Second {
		t.Errorf("expect sleep 4 second but got %v", sleepDur)
	}

	// don't need sleep
	var sleepDur2 time.Duration
	s2 := SleeperOptimized{
		d: 5 * time.Second,
		sleepFunc: func(duration time.Duration) {
			sleepDur = duration
		},
	}

	s2.SleepUntil(context.Background(), 10*time.Second)
	if sleepDur2 != 0 {
		t.Errorf("expect don't sleep but got %v", sleepDur)
	}
}
