package main

import (
	"fmt"
	"net"
)

// Fix-me!
func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		return
	}

	fmt.Println("Launching server in 0.0.0.0:3000...")

	conn, err := ln.Accept()
	if err != nil {
		return
	}

	defer conn.Close()

	for {
		raw := make([]byte, 1024)

		length, err := conn.Read(raw)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Read", length, "bytes.")

		raw = raw[:length]

		length, err = conn.Write(raw)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Write", length, "bytes.")
	}
}
