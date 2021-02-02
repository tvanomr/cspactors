package cspactors

import (
	"github.com/tvanomr/cspactors/lflist"
)

type Actor struct {
	queue    lflist.List
	incoming <-chan interface{}
}
