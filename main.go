package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Printf("Was not able to connect to server: %s", err)
	}
	defer conn.Close()

	var wg sync.WaitGroup

	wg.Add(2)
	go Writer(conn, &wg)
	go Reader(conn, &wg)

	wg.Wait()
}

func Reader(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	scaner := bufio.NewScanner(conn)
	for scaner.Scan() {
		line := scaner.Text()
		fmt.Println(line)
	}
}

func Writer(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	scaner := bufio.NewScanner(os.Stdin)

	for scaner.Scan() {
		line := scaner.Text()
		conn.Write([]byte(line + "\n"))
	}
}
