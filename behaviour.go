package cspactors

import (
	"context"
	"github.com/goccy/go-reflect"
)

type CommandProcessor func(ctx context.Context, cmd interface{}) Reply

type Commands map[uintptr]CommandProcessor

func (c *Commands) AddCommand(arg interface{}, processor CommandProcessor) *Commands {
	if *c == nil {
		*c = make(Commands, 1)
	}
	(*c)[reflect.TypeID(arg)] = processor
	return c
}

func (c *Commands) getCommandProcessor(arg interface{}) CommandProcessor {
	return (*c)[reflect.TypeID(arg)]
}

type MessageProcessor func(msg interface{})

type Messages map[uintptr]MessageProcessor

func (m *Messages) AddMessage(arg interface{}, processor MessageProcessor) *Messages {
	if *m == nil {
		*m = make(Messages, 1)
	}
	(*m)[reflect.TypeID(arg)] = processor
	return m
}

func (m *Messages) getMessageProcessor(arg interface{}) MessageProcessor {
	return (*m)[reflect.TypeID(arg)]
}

type Behaviour struct {
	Commands
	Messages
}
