package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleep() {
	time.Sleep(
		time.Duration(rand.Intn(3000)) * time.Millisecond,
	)
}

func producter(ch chan<- *Item, name string) {
	for {
		sleep()

		n := rand.Intn(100)

		fmt.Printf("Channel %s -> %d\n", name, n)
		item := &Item{name: name, value: n}
		ch <- item
	}
}

func consumer(ch <-chan *Item) {
	for item := range ch {
		fmt.Printf("<- %d, from %s\n", item.value, item.name)
	}
}

func fanIn(chA, chB <-chan *Item, chC chan<- *Item) {
	var item *Item
	for {
		select {
		case item = <-chA:
			chC <- item
		case item = <-chB:
			chC <- item
		}
	}
}

type Item struct {
	name  string
	value int
}

func main() {
	chA := make(chan *Item)
	chB := make(chan *Item)
	chC := make(chan *Item)

	go producter(chA, "A")
	go producter(chB, "B")
	go consumer(chC)

	fanIn(chA, chB, chC)
}
