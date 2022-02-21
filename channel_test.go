package belajargolanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Andry Halomoan Ompusunggu"
	}()

	data := <-channel
	fmt.Println(data)
	close(channel)

}

func GiveMeRespond(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Andry Halomoan Ompusunggu"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeRespond(channel)

	data := <-channel
	fmt.Println(data)
	close(channel)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Andry Halomoan Ompusunggu"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}
