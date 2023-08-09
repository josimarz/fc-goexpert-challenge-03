package web

import (
	"encoding/json"
	"net/http"

	"github.com/josimarz/fc-goexpert-challenge-03/internal/entity"
	"github.com/josimarz/fc-goexpert-challenge-03/internal/usecase"
	"github.com/josimarz/fc-goexpert-challenge-03/pkg/events"
)

type WebOrderHandler struct {
	repository entity.OrderRepository
	event      events.Event
	dispatcher *events.EventDispatcher
}

func NewWebOrderHandler(repository entity.OrderRepository, event events.Event, dispatcher *events.EventDispatcher) *WebOrderHandler {
	return &WebOrderHandler{
		repository,
		event,
		dispatcher,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateOrderInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uc := usecase.NewCreateOrderUseCase(h.repository, h.event, h.dispatcher)
	output, err := uc.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *WebOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	uc := usecase.NewListOrdersUseCase(h.repository)
	output, err := uc.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
