package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()
	
	conn, err := net.DialTimeout("tcp", "0.0.0.0:8080", *timeout)
	if err != nil {
		log.Println(err)
		return
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<- sig
		
		conn.Close()
		os.Exit(0)
			
	}()
    

	buff := make([]byte, 128)
	for {
		io.Copy(conn, os.Stdin)
		
		n, _ := conn.Read(buff)
		if string(buff[0:n]) == os.Interrupt.String() {
			sig <- os.Interrupt
			fmt.Println("msg")
		}
	}

}