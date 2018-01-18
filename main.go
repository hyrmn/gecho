package main

import (
	"flag"
	"io"
	"log"
	"net"
)

func main() {
	port := flag.String("p", "7", "The default port to listen on")
	flag.Parse()

	log.Printf("listening on port %v", *port)

	ln, err := net.Listen("tcp", ":"+*port)

	if err != nil {
		log.Fatalf("listen error, err=%s", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("accept error, err=%s", err)
			return
		}

		go handleConn(conn)
		log.Printf("connection accepted %v", conn.RemoteAddr())
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	_, err := io.Copy(conn, conn)
	if err != nil {
		log.Fatalf("copy error, err=%s", err)
	}
}
