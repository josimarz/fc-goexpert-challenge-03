package events

import (
	"fmt"
	"sync"
)

type EventDispatcher struct {
	handlers map[string][]EventHandler
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (ed *EventDispatcher) Register(name string, handler EventHandler) error {
	if _, ok := ed.handlers[name]; ok {
		for _, h := range ed.handlers[name] {
			if h == handler {
				return fmt.Errorf("handler for event %s already registered", name)
			}
		}
	}
	ed.handlers[name] = append(ed.handlers[name], handler)
	return nil
}

func (ed *EventDispatcher) Remove(name string, handler EventHandler) {
	if _, ok := ed.handlers[name]; ok {
		for i, h := range ed.handlers[name] {
			if h == handler {
				ed.handlers[name] = append(ed.handlers[name][:i], ed.handlers[name][i+1:]...)
				return
			}
		}
	}
}

func (ed *EventDispatcher) Dispatch(event Event) error {
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

func (ed *EventDispatcher) Has(name string, handler EventHandler) bool {
	if _, ok := ed.handlers[name]; ok {
		for _, h := range ed.handlers[name] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandler)
}
