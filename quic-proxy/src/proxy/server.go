package proxy

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"net"
	"os"

	"github.com/quic-go/quic-go"
)

func RunServer(listenTCP, quicAddr string) error {
	// Start QUIC listener
	tlsConf := generateTLSConfig()
	quicListener, err := quic.ListenAddr(quicAddr, tlsConf, nil)
	if err != nil {
		return err
	}
	log.Printf("QUIC server listening on %s", quicAddr)

	// Accept QUIC session from client
	go func() {
		for {
			// sess, err := quicListener.Accept(context.Background())
			_, err := quicListener.Accept(context.Background())
			if err != nil {
				log.Println("QUIC accept error:", err)
				continue
			}
			//go handleQUICSession(sess)
		}
	}()

	// Listen for incoming TCP connections
	ln, err := net.Listen("tcp", listenTCP)
	if err != nil {
		return err
	}
	log.Printf("TCP server listening on %s", listenTCP)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("TCP accept error:", err)
			continue
		}
		go handleTCPConn(conn, quicListener)
	}
}

func handleTCPConn(tcpConn net.Conn, quicListener *quic.Listener) {
	defer tcpConn.Close()
	// Accept a QUIC session (wait for client)
	sess, err := (*quicListener).Accept(context.Background())
	if err != nil {
		log.Println("QUIC session accept error:", err)
		return
	}
	stream, err := sess.OpenStreamSync(context.Background())
	if err != nil {
		log.Println("QUIC stream open error:", err)
		return
	}
	defer stream.Close()
	// Bidirectional copy
	log.Println("Handling TCP connection:", tcpConn.RemoteAddr())
	go io.Copy(stream, tcpConn)
	io.Copy(tcpConn, stream)
}

/*func handleQUICSession(sess quic.Connection) {
	defer sess.CloseWithError(0, "session closed")
	for {
		stream, err := sess.AcceptStream(context.Background())
		if err != nil {
			log.Println("Error accepting QUIC stream:", err)
			return
		}
		// For this simple proxy, just close the stream immediately.
		// In a more advanced implementation, you could handle incoming streams here.
		stream.Close()
	}
}*/

func generateTLSConfig() *tls.Config {
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{"quic-proxy"},
	}
}

// Self-signed cert for demo purposes
var certPEM, _ = os.ReadFile("cert.pem")
var keyPEM, _ = os.ReadFile("key.pem")
