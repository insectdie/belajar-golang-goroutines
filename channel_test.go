package belajargolanggoroutines

import (
	"fmt"
	"strconv"
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

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 1)
	defer close(channel)

	channel <- "Andry"

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan Ke - " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima Data", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRespond(channel1)
	go GiveMeRespond(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRespond(channel1)
	go GiveMeRespond(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
		default:
			fmt.Println("menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
