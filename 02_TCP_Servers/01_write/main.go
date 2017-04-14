package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	//defines li to a listener interface
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	//defer the close of the listener
	defer li.Close()
	//loop forever until someone calls in and accepts connection
	for {
		//conn is both reader/writer interfaces.
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}
		//writes string to writer interface
		io.WriteString(conn, "\nHello from TCP serve\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		//have to close connection.
		conn.Close()
	}
}
