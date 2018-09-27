package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var muxMap map[string]interface{}

func init() {
	muxMap = map[string]interface{}{
		"/":                      endpointIndex,
		"/api/trevor":            endpoint1,
		"/api/scheurer":          endpoint2,
		"/api/golang/is/amazing": endpoint3,
	}
}

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
	request(conn)
}

func request(conn net.Conn) {
	i := 0

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[1]
			for k, v := range muxMap {
				if k == m {
					v.(func(net.Conn, string))(conn, m)
				}
			}
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func endpointIndex(conn net.Conn, s string) {

	body := fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>%v</strong></body></html>`, s)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func endpoint1(conn net.Conn, s string) {

	body := fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>%v</strong></body></html>`, s)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func endpoint2(conn net.Conn, s string) {

	body := fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>%v</strong></body></html>`, s)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func endpoint3(conn net.Conn, s string) {

	body := fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>%v</strong></body></html>`, s)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
