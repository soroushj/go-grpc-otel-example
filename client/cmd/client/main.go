package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/soroushj/go-grpc-otel-example/client"
	"github.com/soroushj/go-grpc-otel-example/jaeger"
	"github.com/soroushj/go-grpc-otel-example/notes"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serviceName = "go-grpc-otel-example/client/cmd/client"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: client server_addr")
	}
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}
	addr := flag.Arg(0)
	ctx := context.Background()
	err := run(ctx, addr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context, addr string) error {
	tp, err := jaeger.TracerProvider(os.Getenv("JAEGER_URL"), serviceName)
	if err != nil {
		return err
	}
	otel.SetTracerProvider(tp)
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		return err
	}
	defer conn.Close()
	nc := notes.NewNotesClient(conn)
	c := client.New(nc)
	r := bufio.NewReader(os.Stdin)
	fmt.Println("enter \\q to quit")
	for {
		fmt.Print("\nid=")
		id, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		id = strings.TrimSpace(id)
		if id == "\\q" {
			return nil
		}
		text, err := c.GetNoteText(ctx, id)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(text)
		}
	}
}
