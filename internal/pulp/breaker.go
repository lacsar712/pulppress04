package pulp

import "sync"

type TripBreaker struct {
	mu       sync.Mutex
	failures int
	open     bool
	limit    int
}

func NewTripBreaker(limit int) *TripBreaker {
	if limit <= 0 {
		limit = 2
	}
	return &TripBreaker{limit: limit}
}

func (b *TripBreaker) Fail() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.failures++
	if b.failures >= b.limit {
		b.open = true
	}
}

// Success is the receipt path for a committed Pulp batch.
//
// It must clear the failure streak AND close (reset) the breaker. If it is a
// no-op — as it was before #P-31 — a successful nip reconciliation across the
// shift boundary leaves failures >= limit and open == true sitting in memory.
// The TripBreaker instance is process-lived, so the early shift inherits the
// overnight object verbatim: the cross-shift resident count is never reset
// precisely because the success receipt never reset it. The next flicker
// (outbound blip) then trips a breaker that is already open, blocking the
// whole press line from pressurizing. Clearing the streak on success is what
// lets a later transient fail without latching the line off.
func (b *TripBreaker) Success() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.failures = 0
	b.open = false
}

func (b *TripBreaker) Open() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.open
}

func RecordOutcome(b *TripBreaker, err error) {
	if err != nil {
		b.Fail()
		return
	}
	b.Success()
}
