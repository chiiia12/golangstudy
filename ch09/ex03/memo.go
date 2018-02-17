package ch09

import (
	"log"
	"fmt"
)

type request struct {
	key      string
	response chan<- result
}
type entry struct {
	res   result
	ready chan struct{}
}
type Memo struct {
	requests chan request
}
type Func func(key string, done chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func, done chan struct{}) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f, done)
	return memo
}

func (memo *Memo) Get(key string, done chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	select {
	case <-done:
		fmt.Println("done has received")
		return res.value, fmt.Errorf("%v request has canceled", key)
	default:
		return res.value, res.err

	}
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func, done chan struct{}) {
	if isCanceled(done) {
		return
	}
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		select {
		case <-done:
			log.Println("done has canceled")
			e = nil
			return
		default:
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, done)
			}
			go e.deliver(req.response)
		}
	}
}

func (e *entry) call(f Func, key string, done chan struct{}) {
	e.res.value, e.res.err = f(key, done)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}

func isCanceled(done chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
