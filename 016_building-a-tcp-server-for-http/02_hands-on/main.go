package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	url := request(conn)

	// write response
	respond(conn, url)
}

func request(conn net.Conn) string {
	i := 0
	var r, m string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			// request line
			m = strings.Fields(ln)[1]
		}
		if i == 1 {
			r = strings.Fields(ln)[1]
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	return r + m
}

func respond(conn net.Conn, s string) {

	body := fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>%v</strong></body></html>`, s)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
