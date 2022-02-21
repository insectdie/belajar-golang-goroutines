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
