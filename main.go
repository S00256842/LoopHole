package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: loophole <receive|send> <port> [message]")
		os.Exit(1)
	}

	if os.Args[1] == "receive" {
		receive(os.Args[2])
	} else if os.Args[1] == "send" {
		send(os.Args[2], os.Args[3])
	} else {
		fmt.Println("Invalid command. Use 'receive' or 'send'.")
	}
}

func receive(port string) {
	addr, _ := net.ResolveUDPAddr("udp", ":"+port)
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, sender, _ := conn.ReadFromUDP(buf)
		fmt.Printf("From %s: %s\n", sender, string(buf[:n]))
	}
}

func send(port string, message string) {
	addr, _ := net.ResolveUDPAddr("udp", "localhost:"+port)
	conn, _ := net.DialUDP("udp", nil, addr)
	conn.Write([]byte(message))
	fmt.Println("Sent message: " + message)
}
