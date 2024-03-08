package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"
	}
}

func pong(c chan string) {
	for {
		c <- "pong"
	}
}

func imprimir(c chan string) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case msg1 := <-c:
			fmt.Println(msg1)
		case msg1 := <-c:
			fmt.Println(msg1)
		}
	}
}

func main() {
	c := make(chan string)

	go ping(c)
	go pong(c)
	go imprimir(c)

	var resposta string
	fmt.Scanln(&resposta)

}
