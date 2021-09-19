package eventsource

import "time"

// Default implementation of an event
type Model struct {
	// the Aggregate ID
	ID string

	// event version number
	Version int

	// eventAt timestamp
	At time.Time
}

// Getters
func (m Model) AggregateID() string {
	return m.ID
}
