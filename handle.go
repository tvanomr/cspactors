package cspactors

import (
	"context"
)

type ActorCompatible interface {
	getActor() *Actor
	MakeBehaviour() Behaviour
}

type Handle struct {
	actor ActorCompatible
	//buffered channel
	incoming    chan<- interface{}
	doneChannel chan struct{}
}

func zombie(msg interface{}) {
	switch value := msg.(type) {
	case *command:
		value.replyWithError(ErrActorDead)
	}
}

func (h *Handle) enqueue(msg interface{}) {
	if h.actor.getActor().queue.Push(msg) {
		select {
		case h.incoming <- nil:
		default:
		}
	} else {
		zombie(msg)
	}
}

func (h *Handle) sendCommand(ctx context.Context, source *Handle, id commandId, cmd interface{}) error {
	select {
	case h.incoming <- &command{ctx, source, id, cmd}:
		return nil
	case <-h.doneChannel:
		return ErrActorDead
	case <-ctx.Done():
		return ErrContextDone
	}
}
