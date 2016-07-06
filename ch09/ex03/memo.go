package memo

import "errors"

type Cancel chan struct{}

func (c Cancel) done() bool {
	select {
	case <-c:
		return true
	default:
		return false
	}
}

// Func is the type of the function to memoize.
type Func func(key string, cancel Cancel) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
	cancel   Cancel
}

type Memo struct{ requests chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, cancel Cancel) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, cancel}
	select {
	case res := <-response:
		return res.value, res.err
	case <-cancel:
		return nil, errors.New("cancelled")
	}
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key, req.cancel) // call f(key)
		}
		go e.deliver(req.response, req.cancel)
	}
}

func (e *entry) call(f Func, key string, cancel Cancel) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key, cancel)
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result, cancel Cancel) {
	// Wait for the ready condition.
	select {
	case <-e.ready:
		// Send the result to the client.
		response <- e.res
	case <-cancel:
	}
}
