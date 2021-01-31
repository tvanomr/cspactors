package lflist

import (
	"sync/atomic"
	"unsafe"
)

type Item struct {
	Data interface{}
	Next *Item
}

type List struct {
	head unsafe.Pointer
}

func (l *List) Push(data interface{}) bool {
	head := atomic.LoadPointer(&l.head)
	if head == unsafe.Pointer(l) {
		return false
	}
	next := &Item{data, (*Item)(head)}
	for !atomic.CompareAndSwapPointer(&l.head, head, unsafe.Pointer(next)) {
		head = atomic.LoadPointer(&l.head)
		if head == unsafe.Pointer(l) {
			return false
		}
		next.Next = (*Item)(head)
	}
	return true
}

func (l *List) TakeHead() *Item {
	head := atomic.LoadPointer(&l.head)
	if head == unsafe.Pointer(l) {
		return nil
	}
	for !atomic.CompareAndSwapPointer(&l.head, head, nil) {
		head = atomic.LoadPointer(&l.head)
		if head == unsafe.Pointer(l) {
			return nil
		}
	}
	return (*Item)(head)
}

func (l *List) TakeHeadAndSeal() *Item {
	head := atomic.LoadPointer(&l.head)
	if head == unsafe.Pointer(l) {
		return nil
	}
	for !atomic.CompareAndSwapPointer(&l.head, head, unsafe.Pointer(l)) {
		head = atomic.LoadPointer(&l.head)
		if head == unsafe.Pointer(l) {
			return nil
		}
	}
	return (*Item)(head)
}
