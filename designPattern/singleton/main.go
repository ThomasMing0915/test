package main

import "sync"

//消息池
type messagePool struct {
	pool *sync.Pool
}

//“饿汉模式”，实例在系统加载的时候就已经完成了初始化

//消息池单例
var msgPool = &messagePool{
	pool: &sync.Pool{New: func() interface{} { return &Message{Count: 0} }},
}

type Message struct {
	Count int
}

//访问消息池单例的唯一方法
func Instance() *messagePool {
	return msgPool
}

//往消息池里添加消息
func (m *messagePool) AddMsg(msg *Message) {
	m.pool.Put(msg)
}

//从消息池里获取消息
func (m *messagePool) GetMsg() *Message {
	return m.pool.Get().(*Message)
}

//懒汉模式
var once = sync.Once{}

var msgPool2 *messagePool

func Instance2() *messagePool {
	once.Do(func() {
		msgPool2 = &messagePool{pool: &sync.Pool{New: func() interface{} { return &Message{Count: 0} }}}
	})
	return msgPool2
}
