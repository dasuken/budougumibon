package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dasuken/budougumibon/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

func run(ctx context.Context) error {
	cfg, _ := config.New()
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d, err: %v \n", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with %s\n", url)

	mux := NewMux()
	s := NewServer(l, mux)

	return s.Run(ctx)
}
