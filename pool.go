package main

import "sync"

type Message struct {
	Text string
}

var p = sync.Pool{
	New: func() any { return new(Message) },
}

func poolExample() {
	v := p.Get().(*Message)
	defer p.Put(v)

	v.Text = "hello guys"
}
