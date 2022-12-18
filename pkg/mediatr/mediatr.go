package mediatr

import (
	"context"
	"fmt"
)

type Mediatr struct {
	requestHandlerRegistry      map[string]RequestHandler
	notificationHandlerRegistry map[string]NotificationHandler
}

type Request interface {
	Topic() string
}

type RequestHandler interface {
	Topic() string
	Handle(ctx context.Context, request Request) (any, error)
}

type NotificationHandler interface {
	Topic() string
	Handle(ctx context.Context, notification Request) error
}

func NewMediatr() *Mediatr {
	return &Mediatr{
		requestHandlerRegistry:      map[string]RequestHandler{},
		notificationHandlerRegistry: map[string]NotificationHandler{},
	}
}

func (m *Mediatr) RegisterRequestHandler(handler RequestHandler) error {
	topic := handler.Topic()

	_, ok := m.requestHandlerRegistry[topic]
	if ok {
		return fmt.Errorf("%w: %s", ErrRequestHandlerAlreadyExists, topic)
	}

	m.requestHandlerRegistry[topic] = handler
	return nil
}

func (m *Mediatr) RegisterNotificationHandler(handler NotificationHandler) error {
	topic := handler.Topic()

	_, ok := m.notificationHandlerRegistry[topic]
	if ok {
		return fmt.Errorf("%w: %s", ErrNotificationHandlerAlreadyExists, topic)
	}

	m.notificationHandlerRegistry[topic] = handler
	return nil
}

func (m *Mediatr) Send(ctx context.Context, request Request) (any, error) {
	topic := request.Topic()

	handler, ok := m.requestHandlerRegistry[topic]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrRequestHandlerNotFound, topic)
	}

	response, err := handler.Handle(ctx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (m *Mediatr) Publish(ctx context.Context, request Request) error {
	topic := request.Topic()

	handler, ok := m.notificationHandlerRegistry[topic]
	if !ok {
		return fmt.Errorf("%w: %s", ErrRequestHandlerNotFound, topic)
	}

	err := handler.Handle(ctx, request)
	if err != nil {
		return err
	}

	return nil
}
