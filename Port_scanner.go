package main

import (
	"fmt"
	"net"
	"sync"
)

func scanPort(address string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
	if err == nil {
		conn.Close()
		fmt.Printf("%d open\n", port)
	}
}

func main() {
	var website string
	fmt.Print("Enter the website to scan: ")
	fmt.Scanln(&website)

	var wg sync.WaitGroup
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go scanPort(website, i, &wg)
	}
	wg.Wait()
}
