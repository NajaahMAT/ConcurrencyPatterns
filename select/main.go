package main

import (
	"fmt"
	"time"
)

func main() {
	//Example 1
	//rpcBlocking()
	//Example 2
	// timeout()
	//Example3
	nonBlocking()

}

func rpcBlocking() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

//output:
// received one
// received two

func timeout() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println("received ", res)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println("received ", res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}

//output:
//received  result 1
//timeout 2

func nonBlocking() {
	messages := make(chan string)
	signals := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		messages <- "Hi"
	}()

	////A non-blocking send works similarly.
	// Here msg cannot be sent to the messages channel, because the channel has no buffer and there is no receiver.
	// Therefore the default case is selected.

	select {
	case msg := <-messages:
		fmt.Println("received message ", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

//output:
//no activity
