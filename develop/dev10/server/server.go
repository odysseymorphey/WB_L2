package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// func handleConnection(ctx context.Context, conn net.Conn) {
// 	for {
// 		select {
// 		case <- ctx.Done():
// 			conn.Write([]byte(os.Interrupt.String()))
// 			conn.Close()
// 		default:
// 			io.Copy(os.Stdout, conn)
// 		}
// 	}
// }

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Println(err)
	}
	log.Println("Server is listening on 0.0.0.0:8080")

	wg := sync.WaitGroup{}
	
	sig := make(chan os.Signal, 1)
	sig2 := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<- sig
		sig2 <- os.Interrupt

		wg.Wait()
		os.Exit(0)
	}()
		
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		wg.Add(1)
		go func(chan os.Signal, net.Conn) {
			for {
				select {
				case <- sig:
					conn.Write([]byte(os.Interrupt.String()))
					conn.Close()
					wg.Done()
					return
				default:
					io.Copy(os.Stdout, conn)
				}
			}
		}(sig2, conn)
	}
}