package main

import (
	"./config_reader"
	"./udp_server"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	config := config_reader.ReadConfigFile("./settings.json")

	request_udp := make(chan string)

	tcp_listener, err := net.Listen("tcp", config["tcp_port"].(string))
	if err != nil {
		log.Fatal(err)
	}

	response := make(chan string)
	if config["address"] == nil {
		fmt.Println("Need an address in settings.json")
	}

	udp_server.CreateServer(config["address"].(string), request_udp)

	// The UDP server
	go func(in, out chan string) {
		for {
			message := <-in

			switch message {
			case "exit":
				out <- "exit"
			default:
				fmt.Printf("%#v\n", strings.Trim(message, "\x00"))
			}
		}
	}(request_udp, response)

	// The TCP server
	go func(listener net.Listener, out chan string) {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(conn)
			conn.Write([]byte("HELLO WORLD\r\n"))
			conn.Close()
		}
	}(tcp_listener, response)

	for {
		message := <-response

		switch message {
		case "exit":
			os.Exit(0)
		}
	}
}
