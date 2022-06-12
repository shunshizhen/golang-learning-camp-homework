package homeworkforsliding

import (
	"errors"
	"sync"
	"time"
)

type timeSlot struct {
	begin time.Time
	count int
}

type SlidingWindowLimiter struct {
	slotDuration time.Duration
	winDuration  time.Duration
	maxReq       int
	windows      []*timeSlot
	mu           sync.Mutex
}

func NewSlidingWindowLimiter(slotDuration, winDuration time.Duration, maxReq int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		slotDuration: slotDuration,
		winDuration:  winDuration,
		maxReq:       maxReq,
	}
}

func (l *SlidingWindowLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	timeoutOffset := -1
	for i, ts := range l.windows {
		if ts.begin.Add(l.winDuration).After(now) {
			break
		}
		timeoutOffset = i
	}
	if timeoutOffset > -1 {
		l.windows = l.windows[timeoutOffset+1:]
	}

	var ok bool
	if l.countReq() < l.maxReq {
		ok = true
	}

	var lastSlot *timeSlot
	if len(l.windows) > 0 {
		lastSlot = l.windows[len(l.windows)-1]
		if lastSlot.begin.Add(l.slotDuration).Before(now) {
			lastSlot = &timeSlot{begin: now, count: 1}
			l.windows = append(l.windows, lastSlot)
		} else {
			lastSlot.count++
		}
	} else {
		lastSlot = &timeSlot{begin: now, count: 1}
		l.windows = append(l.windows, lastSlot)
	}

	return ok
}

func (l *SlidingWindowLimiter) countReq() int {
	var count int
	for _, ts := range l.windows {
		count += ts.count
	}
	return count
}

func (l *SlidingWindowLimiter) Run(fn func()) error {
	if !l.Allow() {
		return errors.New("error:")
	}
	fn()
	return nil
}
