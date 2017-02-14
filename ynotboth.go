package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/akresling/ynotboth/pb"
	"github.com/akresling/ynotboth/service"
)

func main() {
	es := service.Example{}
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		port := ":8080"
		fmt.Println("starting on 8080")
		errc <- http.ListenAndServe(port, service.Router(es))
	}()

	go func() {
		listen, _ := net.Listen("tcp", ":3333")
		grpcServer := grpc.NewServer()
		pb.RegisterExampleServer(grpcServer, es)
		errc <- grpcServer.Serve(listen)
	}()

	// Run server until termination signal
	log.Printf("exit %s \n", <-errc)
}
