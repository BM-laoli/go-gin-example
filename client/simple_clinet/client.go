package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/BM-laoli/go-gin-example/proto"
)

const PORT = "9001"

func main() {
	cert, err := tls.LoadX509KeyPair("../../conf/client_cert.pem", "../../conf/client_key.pem")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../../conf/ca_cert.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "x.test.example.com",
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(":"+PORT, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())
}

// openssl req -new -nodes -key hello.key -out hello.csr -days 3650 -subj "/C=/ST=/L=/OU=/O=/CN=go-gin" -config ./openssl.cnf -extensions v3_req

// openssl x509 -req -days 3650 -in hello.csr -out hello.pem -CA ca.pem -CAkey server.key -CAcreateserial -extfile ./openssl.cnf -extensions v3_req
