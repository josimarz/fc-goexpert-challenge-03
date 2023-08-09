package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/josimarz/fc-goexpert-challenge-03/pkg/events"
	"github.com/streadway/amqp"
)

type OrderCreatedHandler struct {
	channel *amqp.Channel
}

func NewOrderCreatedHandler(channel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{channel}
}

func (h *OrderCreatedHandler) Handle(event events.Event, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order created: %v", event.GetPayload())
	body, _ := json.Marshal(event.GetPayload())
	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}
	h.channel.Publish(
		"amq.direct",
		"",
		false,
		false,
		msg,
	)
}
