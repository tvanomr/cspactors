package cspactors

import ()

type actorError int

const (
	ErrActorDead actorError = iota
	ErrContextDone
)

var errorMessages = map[actorError]string{
	ErrActorDead:   "Actor is dead",
	ErrContextDone: "Context finished before command was processed"}

func (e actorError) Error() string {
	return errorMessages[e]
}
