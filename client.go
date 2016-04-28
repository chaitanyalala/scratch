package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// First Argument is server IP:Port e.g. 127.0.0.1:42323,
// third FS path on remote server
func parseArgs(args []string) (addr string, path string) {
	length := len(args)
	if length < 3 {
		log.Fatal("Please pass arguments e.g. 127.0.0.1:42323 /home/testfile.txt")
	}
	if args[1] != "" {
		addr = args[1]
	}
	if args[2] != "" {
		path = args[2]
	}
	return addr, path
}

func dialTcp(remoteIpPort string, path string) {
	conn, err := net.Dial("tcp", remoteIpPort)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(conn, path)
	conn.Close()
}

func main() {
	addr, path := parseArgs(os.Args)
	for i := 0; i < 1000*1000; i++ {
		go dialTcp(addr, path)
	}
}
