package main

import (
	"fmt"
	"sync"
)

type Header struct {
	SrcAddr  string
	SrcPort  uint64
	DestAddr string
	DestPort uint64
	Items    map[string]string
}

type Body struct {
	Items []string
}

type Message struct {
	Header *Header
	Body   *Body
}

var message = Message{
	Header: &Header{
		SrcAddr:  "127.0.0.1",
		SrcPort:  1234,
		DestAddr: "127.0.0.1",
		DestPort: 8000,
		Items:    make(map[string]string),
	},
	Body: &Body{
		Items: make([]string, 0),
	},
}

func main() {
	message.Header.Items["contents"] = "application/json"
	message.Body.Items = append(message.Body.Items, "record1")
	message.Body.Items = append(message.Body.Items, "record2")

	msg := Builder().WithSrcAddr("127.0.0.1").
		WithSrcPort(1234).
		WithDestAddr("127.0.0.1").
		WithDestPort(8000).
		WithHeaderItem("contents", "application/json").
		WithBodyItem("record1").
		WithBodyItem("record2").
		Build()

	if msg.Header.SrcAddr != "127.0.0.1" {
		fmt.Printf("expect is 127.0.0.1 actual is %s\n", msg.Header.SrcAddr)
	}

	if msg.Body.Items[0] != "record1" {
		fmt.Printf("expect is record1 actual is %s\n", msg.Body.Items[0])
	}

}

//Message对象的Builder对象
type builder struct {
	once *sync.Once
	msg  *Message
}

func Builder() *builder {
	return &builder{
		once: &sync.Once{},
		msg: &Message{
			Header: &Header{},
			Body:   &Body{},
		},
	}
}

func (b *builder) WithSrcAddr(srcAddr string) *builder {
	b.msg.Header.SrcAddr = srcAddr
	return b
}

func (b *builder) WithSrcPort(srcPort uint64) *builder {
	b.msg.Header.SrcPort = srcPort
	return b
}

func (b *builder) WithDestAddr(destAddr string) *builder {
	b.msg.Header.DestAddr = destAddr
	return b
}

func (b *builder) WithDestPort(destPort uint64) *builder {
	b.msg.Header.DestPort = destPort
	return b
}

func (b *builder) WithHeaderItem(key, value string) *builder {
	b.once.Do(func() {
		b.msg.Header.Items = make(map[string]string)
	})
	b.msg.Header.Items[key] = value
	return b
}

func (b *builder) WithBodyItem(record string) *builder {
	b.msg.Body.Items = append(b.msg.Body.Items, record)
	return b
}

func (b *builder) Build() *Message {
	return b.msg
}
