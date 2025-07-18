package proxy

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"net"

	"github.com/quic-go/quic-go"
)

func RunClient(quicAddr, targetTCP string) error {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-proxy"},
	}
	sess, err := quic.DialAddr(context.Background(), quicAddr, tlsConf, nil)
	if err != nil {
		return err
	}
	log.Printf("Connected to QUIC server at %s", quicAddr)

	for {
		stream, err := sess.AcceptStream(context.Background())
		if err != nil {
			log.Println("QUIC stream accept error:", err)
			continue
		}
		go handleQUICStream(stream, targetTCP)
	}
}

func handleQUICStream(stream *quic.Stream, targetTCP string) {
	defer stream.Close()
	tcpConn, err := net.Dial("tcp", targetTCP)
	if err != nil {
		log.Println("TCP dial error:", err)
		return
	}
	defer tcpConn.Close()
	// Bidirectional copy
	log.Println(stream)
	go io.Copy(tcpConn, stream)
	io.Copy(stream, tcpConn)
}
