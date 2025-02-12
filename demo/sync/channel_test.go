package sync

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(T *testing.T) {
	ch := make(chan string, 4)
	go func() {
		str := <-ch
		fmt.Println("1", str)
	}()
	go func() {
		str := <-ch
		fmt.Println("2", str)
	}()
	go func() {
		str := <-ch
		fmt.Println("3", str)
	}()

	ch <- "Hello"
	ch <- "Hello,World!"

}

func TestPubAndSub(T *testing.T) {
	c1 := &Consumer{ch: make(chan string, 1)}
	c2 := &Consumer{ch: make(chan string, 1)}
	b := &Broker{consumers: make([]*Consumer, 0, 10)}
	b.Subscribe(c1)
	b.Subscribe(c2)
	b.Publish("Hello")
	fmt.Println(<-c1.ch)
	fmt.Println(<-c2.ch)
}

type Broker struct {
	consumers []*Consumer
}

func (b *Broker) Publish(msg string) {
	for _, c := range b.consumers {
		c.ch <- msg
	}
}

func (b *Broker) Subscribe(c *Consumer) {
	b.consumers = append(b.consumers, c)
}

type Consumer struct {
	ch chan string
}

//-----------------------------------------

type Broker1 struct {
	ch        chan string
	consumers []func(s string)
}

func (b *Broker1) Publish(msg string) {
	b.ch <- msg

}

func (b *Broker1) Subscribe(consumer func(s string)) {
	b.consumers = append(b.consumers, consumer)
}

func (b *Broker1) Start() {
	go func() {
		s := <-b.ch
		for _, c := range b.consumers {
			c(s)
		}
	}()
}

func NewBroker1() *Broker1 {
	b := &Broker1{
		ch:        make(chan string, 10),
		consumers: make([]func(s string), 0, 10),
	}
	b.Start()
	return b
}

func TestPubAndSub1(T *testing.T) {
	broker := NewBroker1()
	// 定义一个订阅者函数
	subscriber1 := func(s string) {
		fmt.Println("Subscriber 1 received:", s)
	}
	subscriber2 := func(s string) {
		fmt.Println("Subscriber 2 received:", s)
	}

	// 订阅消息
	broker.Subscribe(subscriber1)
	broker.Subscribe(subscriber2)
	// 发布消息
	broker.Publish("Hello, World!")
	time.Sleep(1 * time.Second)
}
