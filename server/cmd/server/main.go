package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/soroushj/go-grpc-otel-example/notes"
	"github.com/soroushj/go-grpc-otel-example/server"
	"github.com/soroushj/go-grpc-otel-example/trcprv"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: server addr")
	}
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}
	addr := flag.Arg(0)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen at %v: %v", addr, err)
	}
	tu := os.Getenv("JAEGER_URL")
	tp, err := trcprv.TracerProvider(tu, "go-grpc-otel-example/server/cmd/server")
	if err != nil {
		log.Fatalf("failed to provide tracer at %v: %v", tu, err)
	}
	otel.SetTracerProvider(tp)
	s := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))
	notes.RegisterNotesServer(s, server.New())
	log.Printf("listening at %v", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
