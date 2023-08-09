package event

import "time"

type OrderCreatedEvent struct {
	name    string
	payload interface{}
}

func NewOrderCreatedEvent() *OrderCreatedEvent {
	return &OrderCreatedEvent{name: "OrderCreated"}
}

func (e *OrderCreatedEvent) GetName() string {
	return e.name
}

func (e *OrderCreatedEvent) GetPayload() interface{} {
	return e.payload
}

func (e *OrderCreatedEvent) SetPayload(payload interface{}) {
	e.payload = payload
}

func (e *OrderCreatedEvent) GetDateTime() time.Time {
	return time.Now()
}
