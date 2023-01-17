package main

import (
	"os"
	"log"
	"net"
	"bufio"
)

func panicError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(bufio.NewReader(conn))

	for scanner.Scan() {
		log.Println(scanner.Text())
	}
}

func main() {
	port := os.Getenv("PORT")

	ln, err := net.Listen("tcp", ":" + port)
	panicError(err)

	for {
		conn, err := ln.Accept()
		panicError(err)

		go handleConnection(conn)
	}
}
