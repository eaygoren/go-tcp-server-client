package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	connect, _ := net.Dial("tcp", ":1993")

	fmt.Print("Say something...\n")

	go func() {
		for {
			message, err := bufio.NewReader(connect).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				fmt.Print(message)
			}
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(connect, text)
	}
}
