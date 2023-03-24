package http

import (
	"context"

	"cpn"
)

type Response struct {
	chin  chan *cpn.M
	chout chan *cpn.M
}

func NewResponse() *Response {
	return &Response{
		chin:  make(chan *cpn.M),
		chout: make(chan *cpn.M),
	}
}

func (p *Response) In() chan<- *cpn.M {
	return p.chin
}

func (p *Response) Out() <-chan *cpn.M {
	return p.chout
}

func (p *Response) Run(_ context.Context) {
	defer close(p.chout)
	for m := range p.chin {
		m.Value().(*RequestContext).Flush()
	}
}
