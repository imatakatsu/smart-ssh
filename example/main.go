package main

import (
	"log"
	"net"

	ssh "github.com/imatakatsu/smart-ssh"
)

func main() {
	srv, err := ssh.Listen(":2222")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := srv.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
	net.Listener
}

func handleClient(conn ssh.Conn) {
	conn.Println("hello, world!")
	conn.Close()
}
