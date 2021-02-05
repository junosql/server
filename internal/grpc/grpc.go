package grpc

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Start runs the GRPC server
func Start(tcpPort uint16) {
	listener := startListener(tcpPort)

	startGrpcServer(listener)
}

func startListener(tcpPort uint16) net.Listener {
	listenAddr := fmt.Sprintf(":%s", fmt.Sprint(tcpPort))
	listener, err := net.Listen("tcp", listenAddr)

	if err != nil {
		panic(err)
	}

	log.Infof("Starting listener on port %s", listenAddr)

	return listener
}

func startGrpcServer(listener net.Listener) {
	log.Info("Starting gRPC interface")

	grpcServer := grpc.NewServer()
	err := grpcServer.Serve(listener)

	if err != nil {
		panic(err)
	}

	log.Info("gRPC interface successfully started")
}
