package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

type options struct {
	port string
}

var opt options

func init() {
	flag.StringVar(&opt.port, "p", os.Getenv("PORT"), "The default port to listen on")
	flag.Parse()

	if opt.port == "" {
		opt.port = "7"
	}
}

func main() {
	log.Printf("listening on port %v", opt.port)

	ln, err := net.Listen("tcp", ":"+opt.port)

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
	bytesCopied, err := io.Copy(conn, conn)
	if err == io.EOF {
		log.Printf("received EOF. client disconnected")
		return
	} else if err != nil {
		log.Fatalf("copy error, err=%s", err)
	}
	log.Printf("sent %v bytes", bytesCopied)
}
