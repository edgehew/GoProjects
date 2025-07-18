package main

import (
	"flag"
	"log"
	"os"

	"quic-proxy/src/proxy"
)

func main() {
	mode := flag.String("mode", "server", "Mode: server or client")
	listenTCP := flag.String("listen-tcp", ":8080", "TCP listen address (server mode)")
	quicAddr := flag.String("quic-addr", "localhost:4242", "QUIC listen/connect address")
	targetTCP := flag.String("target-tcp", "localhost:80", "Target TCP address (client mode)")
	flag.Parse()

	switch *mode {
	case "server":
		log.Printf("Starting in server mode: TCP %s <-> QUIC %s\n", *listenTCP, *quicAddr)
		if err := proxy.RunServer(*listenTCP, *quicAddr); err != nil {
			log.Fatal(err)
		}
	case "client":
		log.Printf("Starting in client mode: QUIC %s <-> TCP %s\n", *quicAddr, *targetTCP)
		if err := proxy.RunClient(*quicAddr, *targetTCP); err != nil {
			log.Fatal(err)
		}
	default:
		log.Println("Unknown mode:", *mode)
		os.Exit(1)
	}
}
