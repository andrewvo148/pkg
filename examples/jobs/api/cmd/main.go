package main

import (
	"andrewvo148/pkg/examples/jobs/api/internal/application"
	"andrewvo148/pkg/examples/jobs/api/internal/domain"

	"context"
	"fmt"
)

//
//type server struct {
//	gen.UnimplementedGreeterServer
//}
//
//func NewServer() *server {
//	return &server{}
//}
//
//func (s *server) SayHello(ctx context.Context, in *gen.HelloRequest) (*gen.HelloReply, error) {
//	return &gen.HelloReply{Message: in.Name + " world"}, nil
//}
//func main() {
//	//r := chi.NewRouter()
//	//
//	//// Mount the FileServer handler to serve static files from the "static" directory
//	//r.Mount("/", http.FileServer(http.FS(web.WebUI)))
//	//
//	//http.ListenAndServe(":8080", r)
//
//	// Create a listener on TCP port
//	lis, err := net.Listen("tcp", ":8080")
//	if err != nil {
//		log.Fatalln("Failed to listen:", err)
//	}
//
//	// Create a gRPC server object
//	s := grpc.NewServer()
//	// Attach the Greeter service to the server
//	gen.RegisterGreeterServer(s, &server{})
//	// Serve gRPC Server
//	log.Println("Serving gRPC on 0.0.0.0:8080")
//	go func() {
//		log.Fatalln(s.Serve(lis))
//	}()
//
//	// Create a client connection to the gRPC server we just started
//	// This is where the gRPC-Gateway proxies the requests
//	gwmux := runtime.NewServeMux()
//	opts := []grpc.DialOption{
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	}
//	err = gen.RegisterGreeterHandlerFromEndpoint(context.Background(), gwmux, ":8080", opts)
//	if err != nil {
//		log.Fatalln("Failed to register gateway:", err)
//	}
//
//	gwServer := &http.Server{
//		Addr:    ":8090",
//		Handler: gwmux,
//	}
//
//	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
//	log.Fatalln(gwServer.ListenAndServe())
//}

type CustomHandler struct {
	// Ha tang.
	app application.Application
}

func (h *CustomHandler) Handle(ctx context.Context, event domain.Event) error {
	fmt.Println(event.Name)
	return nil
}

func main() {
	de := domain.NewEventDispatcher()
	fmt.Println(de)
	handler := CustomHandler{}

	err := de.Subscribe(handler, "event1")
	if err != nil {
		return
	}

	de.Publish(context.TODO(), domain.Event{Name: "event1"})
}
