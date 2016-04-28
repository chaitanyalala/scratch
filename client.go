package main

import (
	"fmt"
	"log"
	"net"
	//"os"
	//"strings"
	//"time"
)

func dialTcp() {
	conn, err := net.Dial("tcp", "localhost:2000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(conn, "/home/clala/Downloads/save.txt")
	conn.Close()
}

func main() {
	for i := 0; i < 1000*1000*10; i++ {
		go dialTcp()
	}
}
