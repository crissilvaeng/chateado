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

	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}

		raw := make([]byte, 1024)

		length, err := conn.Read(raw)
		if err != nil {
			fmt.Println("Read error: ", err)
			return
		}

		fmt.Println("Read", length, "bytes.")

		raw = []byte(`{
                "id": "0",
                "msgNr": 54,
                "data": [
                    {"src":"maria","data":"oi!"},
                    {"src":"maria","data":"kd vc?"}
                ]
            }`)

		length, err = conn.Write(raw)
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}

		fmt.Println("Write", length, "bytes.")
	}
}
