package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Open a TCP socket on port 80
	conn, err := net.Dial("tcp", "sarathyweb.com:80")
	if err != nil {
		log.Fatal(err)
	}
	// Send a HTTP GET request
	// CRLF = Carriage Return + Line Feed = \r\n
	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: sarathyweb.com\r\nConnection: close\r\n\r\n")
	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// the respons for a HTTP request is a stream.
	// So we cannot use println to print the response
	// The ReadString function reads from the buffer until it founds a new line character
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// Prints the status
	log.Println(status)
}
