package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
)

func main() {
	conn, err := tls.Dial("tcp", "sarathyweb.com:443", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: sarathyweb.com\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	log.Println(status)
}
