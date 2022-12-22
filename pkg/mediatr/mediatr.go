package mediatr

import (
	"context"
	"fmt"
	"reflect"

	"github.com/ahmetb/go-linq/v3"
	"github.com/ssengalanto/potato-project/pkg/gg"
)

type Mediatr struct {
	requestHandlerRegistry      map[string]RequestHandler
	notificationHandlerRegistry map[string]NotificationHandler
	pipelineBehaviourRegistry   []PipelineBehavior
}

type RequestHandlerFunc func() (any, error)

type RequestHandler interface {
	Name() string
	Handle(ctx context.Context, request any) (any, error)
}

type NotificationHandler interface {
	Name() string
	Handle(ctx context.Context, notification any) error
}

type PipelineBehavior interface {
	Handle(ctx context.Context, request any, next RequestHandlerFunc) (any, error)
}

// New creates a new mediatr instance.
func New() *Mediatr {
	return &Mediatr{
		requestHandlerRegistry:      map[string]RequestHandler{},
		notificationHandlerRegistry: map[string]NotificationHandler{},
		pipelineBehaviourRegistry:   []PipelineBehavior{},
	}
}

// RegisterRequestHandler register a request handler in the registry.
func (m *Mediatr) RegisterRequestHandler(handler RequestHandler) error {
	hn := handler.Name()

	_, ok := m.requestHandlerRegistry[hn]
	if ok {
		return fmt.Errorf("%w: %s", ErrRequestHandlerAlreadyExists, hn)
	}

	m.requestHandlerRegistry[hn] = handler
	return nil
}

// RegisterNotificationHandler register a notification handler in the registry.
func (m *Mediatr) RegisterNotificationHandler(handler NotificationHandler) error {
	hn := handler.Name()

	_, ok := m.notificationHandlerRegistry[hn]
	if ok {
		return fmt.Errorf("%w: %s", ErrNotificationHandlerAlreadyExists, hn)
	}

	m.notificationHandlerRegistry[hn] = handler
	return nil
}

// RegisterPipelineBehaviour register a behaviour in the registry.
func (m *Mediatr) RegisterPipelineBehaviour(behaviour PipelineBehavior) error {
	behaviourType := reflect.TypeOf(behaviour)

	exists := m.existsPipeType(behaviourType)
	if exists {
		return fmt.Errorf("%w: %s", ErrPipelineBehaviourAlreadyExists, behaviourType)
	}

	m.pipelineBehaviourRegistry = gg.Prepend(m.pipelineBehaviourRegistry, behaviour)
	return nil
}

// Send sends the request to its corresponding request handler.
func (m *Mediatr) Send(ctx context.Context, request any) (any, error) {
	rt := reflect.TypeOf(request).String()

	handler, ok := m.requestHandlerRegistry[rt]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrRequestHandlerNotFound, rt)
	}

	if len(m.pipelineBehaviourRegistry) > 0 {
		var lastHandler RequestHandlerFunc = func() (any, error) {
			return handler.Handle(ctx, request)
		}

		aggregateResult := linq.From(m.pipelineBehaviourRegistry).AggregateWithSeedT(
			lastHandler,
			func(next RequestHandlerFunc, pipe PipelineBehavior) RequestHandlerFunc {
				pipeValue := pipe
				nexValue := next

				var handlerFunc RequestHandlerFunc = func() (any, error) {
					return pipeValue.Handle(ctx, request, nexValue)
				}

				return handlerFunc
			})

		v := aggregateResult.(RequestHandlerFunc) //nolint:errcheck //unnecessary
		response, err := v()
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	response, err := handler.Handle(ctx, request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Publish publishes the notification event to its corresponding notification handler.
func (m *Mediatr) Publish(ctx context.Context, request any) error {
	rt := reflect.TypeOf(request).String()

	handler, ok := m.notificationHandlerRegistry[rt]
	if !ok {
		return fmt.Errorf("%w: %s", ErrRequestHandlerNotFound, rt)
	}

	err := handler.Handle(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

// existsPipeType checks if a pipeline behaviour exists in the registry.
func (m *Mediatr) existsPipeType(p reflect.Type) bool {
	for _, pipe := range m.pipelineBehaviourRegistry {
		if reflect.TypeOf(pipe) == p {
			return true
		}
	}

	return false
}
