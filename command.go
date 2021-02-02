package cspactors

import (
	"context"
)

type message struct {
	source Handle
	msg    interface{}
}

type commandId int64

type command struct {
	ctx       context.Context
	source    *Handle
	commandId commandId
	cmd       interface{}
}

func (c *command) reply(data interface{}) {
	c.source.enqueue(&reply{c.commandId, data})
}

func (c *command) replyWithError(err error) {
	c.source.enqueue(&replyError{c.commandId, err})
}

type reply struct {
	commandId commandId
	reply     interface{}
}

type replyError struct {
	commandId  commandId
	replyError error
}
