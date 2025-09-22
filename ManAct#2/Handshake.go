package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go client(ch, 100)
	go server(ch, 300)

	time.Sleep(20 * time.Second)
}

func server(channol chan int, sequence int) {

	ack := <-channol //ACK: Receiving clients sequence
	fmt.Println("Server: Received sequence", ack, "from client")
	ack = ack + 1 //plus 1
	time.Sleep(1 * time.Second)

	fmt.Println("Server: Sending acknowledgement", ack, "to client")
	channol <- ack //Sending acknowledgement to client

	fmt.Println("Server: Sending sequence:", sequence, "to client")
	channol <- sequence //Sending sequence number to client
	time.Sleep(1 * time.Second)

	oldack := ack
	ack = <-channol
	syn := <-channol

	if ack == oldack && syn == sequence+1 {
		fmt.Println("Connection established")
	} else {
		fmt.Println("Failed to connect")
	}
}

func client(channol chan int, seq int) {

	channol <- seq
	fmt.Println("SYN: Client is sending seq number:", seq, "to the server")
	time.Sleep(1 * time.Second)

	var ack = <-channol //receiving acknowledgement from server
	var syn = <-channol //receiving server's seq number

	fmt.Println("client recieved ACK:", ack, "from ther server, and SYN:", syn)

	if ack == seq+1 {
		fmt.Println("the numbers match")
		fmt.Println("client is sending altered seq number", seq+1, "and ACK", syn+1)
		channol <- seq + 1
		channol <- syn + 1
	}

}
