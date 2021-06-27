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

func producter(ch chan<- *Item) {
	for {
		sleep()

		n := rand.Intn(100)

		var item Item
		fmt.Printf("Producter %s -> %d\n", "A", n)
		if n <= 50 {
			item = Item{name: "BB", value: n}
		} else {
			item = Item{name: "CC", value: n}
		}
		ch <- &item
	}
}

func consumer(ch <-chan *Item, name string) {
	for item := range ch {
		fmt.Printf("Consumer %s <- Item(%s, %d)\n", name, item.name, item.value)
	}
}

func fanout(chA <-chan *Item, chB, chC chan<- *Item) {
	for item := range chA {
		if item.value <= 50 {
			chB <- item
		} else {
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

	go producter(chA)
	go consumer(chB, "B")
	go consumer(chC, "C")

	fanout(chA, chB, chC)
}
