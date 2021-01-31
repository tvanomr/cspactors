package cspactors

import ()

type ResponsePromise struct{}

type ReplyVisitor interface {
	Reply(interface{})
	Promise(*ResponsePromise)
}

//reply value or response promise (the way to implement multiple return types in go, runtime stuff obviously)

type Reply interface {
	Visit(ReplyVisitor)
}
