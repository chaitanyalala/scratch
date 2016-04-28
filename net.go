package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
func ReadFile(fileName string) (string, bool) {
	data := make([]byte, 16, 4096)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var ret string
	for {
		count, err := file.Read(data)
		ret += string(data[:count])
		if err != nil {
			return ret, true
		}
	}
	return ret, false
}
*/

// Handle the connection in a new goroutine.
// The loop then returns to accepting, so that
// multiple connections may be served concurrently.
func handleConn(c *net.TCPConn) {
	defer c.Close()
	// Take incoming file name.
	buf := make([]byte, 8, 8)
	var fileName string
	for {
		n, ok := c.Read(buf)
		fileName += fmt.Sprint(string(buf[:n]))
		if ok != nil {
			break
		}
	}
	/*
		d, ok := ReadFile(strings.TrimSpace(fileName))
		if ok == true {
			c.Write([]byte(d))
		}
	*/
	file, err := os.Open(strings.TrimSpace(fileName))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	written, err := c.ReadFrom(file)
	if err != nil {
		fmt.Println("err %v, written %d", err, written)
	}
}

func printConn(cnt int64, start time.Time, old int64) (time.Time, int64) {
	timeNow := time.Now()
	diff := timeNow.Sub(start)
	if diff > 0 {
		fmt.Println(1000 * 1000 * 1000 * (cnt - old) / diff.Nanoseconds())
	}
	old = cnt
	start = time.Now()

	return start, old
}

func parseArgs(args []string) net.TCPAddr {
	var addr net.TCPAddr
	length := len(args)
	if length >= 2 {
		if args[1] != "" {
			addr.IP = net.ParseIP(args[1])
		}
		if length > 2 {
			if args[2] != "" && args[2] != "0" {
				addr.Port, _ = strconv.Atoi(args[2])
			}
		}
	}
	return addr
}

// Args as IP and port to listen on
func main() {
	const listenPort = 2000
	const protocol = "tcp"
	var addr net.TCPAddr
	var cnt, old int64

	addr = parseArgs(os.Args)
	if addr.Port == 0 {
		addr.Port = listenPort
	}
	l, err := net.ListenTCP(protocol, &addr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	start := time.Now()
	for {
		// Wait for a connection.
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		cnt++
		go handleConn(conn)
		start, old = printConn(cnt, start, old)
	}
}
