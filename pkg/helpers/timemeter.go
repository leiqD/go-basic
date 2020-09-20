package helpers

import (
	"time"
)

type TimeMeter interface {
	StepTUsed() time.Duration
}

type timeMeter struct {
	timestamp time.Time
}

func NewTimeMeter() TimeMeter {
	return &timeMeter{
		timestamp: time.Now(),
	}
}

func (t *timeMeter) StepTUsed() time.Duration {
	now := time.Now()
	dur := now.Sub(t.timestamp)
	t.timestamp = now
	return dur
}
