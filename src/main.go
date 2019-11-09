package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/soheilhy/cmux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	kncloudevents "github.com/celsosantos/http-grpc-mux/kncloudevents"
	cloudevents "github.com/cloudevents/sdk-go"

	pb "github.com/celsosantos/http-grpc-mux/api"
)

var (
	ctx = context.Background()

	// Make Channel to receive end-result
	htmlChannel = make(chan string)
)

type htmlService struct{}

func (c *htmlService) Render(ctx context.Context, in *pb.HtmlRequest) (*pb.HtmlResponse, error) {

	contentBytes, err := ioutil.ReadFile("/var/templates/" + os.Getenv("TEMPLATE") + "-template.html")
	if err != nil {
		log.Println("Error: %s", err)
	}

	htmlTemplate := string(contentBytes)

	return &pb.HtmlResponse{Status: 200, Document: htmlTemplate}, nil
}

func render(event cloudevents.Event) {

	log.Println("GOT AN EVENT!!! YAY!!!")
	if event.Data == nil {
		log.Fatalf("Missing event Rendered document.")
	}

	htmlChannel <- "Hello Channel!"
}

func main() {

	port, err := strconv.Atoi(os.Getenv("LISTEN_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create multiplexer and listener types
	mux := cmux.New(lis)
	grpcLis := mux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpLis := mux.Match(cmux.HTTP1Fast())

	// *************
	// gRPC
	// *************

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// create server
	pb.RegisterHtmlServiceServer(grpcServer, &htmlService{})

	//reflection required for Gloo Discovery
	reflection.Register(grpcServer)

	// *************
	// HTTP
	// *************

	// Declare new CloudEvents Receiver
	c, err := kncloudevents.NewDefaultClient(httpLis)
	if err != nil {
		log.Fatal("Failed to create client, ", err)
	}

	defer close(htmlChannel)

	// *************
	// Start Listeners
	// *************

	// start gRPC server
	go func() {
		if err := grpcServer.Serve(grpcLis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	// start HTTP server
	go func() {
		log.Fatal(c.StartReceiver(ctx, render))
	}()

	mux.Serve()
}
