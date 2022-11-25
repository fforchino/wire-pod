package main

import (
	"crypto/x509"
	"fmt"
	tokenpb "github.com/digital-dream-labs/api/go/tokenpb"
	grpcserver "github.com/digital-dream-labs/hugh/grpc/server"
	"github.com/digital-dream-labs/hugh/log"
	"io/ioutil"
	"token/pkg/server"
)

func main() {
	log.SetJSONFormat("2006-01-02 15:04:05")
	startServer()
}

func startServer() {
	certPool := x509.NewCertPool()
	var cert, _ = ioutil.ReadFile("../certs/cert.crt")
	if !certPool.AppendCertsFromPEM(cert) {
		log.Fatal("failed to add server CA's certificate")
	}

	srv, err := grpcserver.New(
		grpcserver.WithPort(445),
		grpcserver.WithViper(),
		grpcserver.WithLogger(log.Base()),
		grpcserver.WithReflectionService(),
		grpcserver.WithInsecureSkipVerify(),

		grpcserver.WithUnaryServerInterceptors(
		//			grpclog.UnaryServerInterceptor(),
		),

		grpcserver.WithStreamServerInterceptors(
		//			grpclog.StreamServerInterceptor(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	s, _ := server.New()

	tokenpb.RegisterTokenServer(srv.Transport(), s)

	srv.Start()
	fmt.Println("\033[33m\033[1mToken Server started successfully!\033[0m")
	<-srv.Notify(grpcserver.Stopped)
}
