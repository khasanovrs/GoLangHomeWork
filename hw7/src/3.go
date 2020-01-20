package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	cancel := make(chan int)
	go func() {
		for {
			b, _ := ioutil.ReadAll(os.Stdin)

			if string(b) == "exit" {
				fmt.Println("exit app")
				cancel <- 1
				return
			}
		}
	}()

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, cancel)
	}
}

func handleConn(c net.Conn, cancel chan int) {
	defer c.Close()

	tick := make(<-chan time.Time)
	tick = time.Tick(1 * time.Second)

	for {
		select {
		case <-tick:
			_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
			if err != nil {
				return
			}
		case <-cancel:
			fmt.Println("Launch canceled!")
			return
		}
	}

}
