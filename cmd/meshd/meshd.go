package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gumball-guardian/meshd"
)

var serverTlsAddr = flag.String("serverTlsAddr", ":4443", "https service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	//systemTeardown()

	list, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for i, iface := range list {
		fmt.Printf("%d name=%s %v\n", i, iface.Name, iface)
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		for j, addr := range addrs {
			fmt.Printf(" %d %v\n", j, addr)
		}
	}

	log.Printf("meshd serving on %s\n", *serverTlsAddr)

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

	log.Fatalln(meshd.ServeTLS(*serverTlsAddr))
}
