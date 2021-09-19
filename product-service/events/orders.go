package events

import (
	"time"

	"github.com/ldenholm/systemsgo/eventsource"
)

type OrderCreated struct {
	eventsource.Model
}

type OrderUpdated struct {
	eventsource.Model
}

type OrderPurchased struct {
	eventsource.Model
}

type ProductOrder struct {
	ID        string
	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time
	State     string
}
