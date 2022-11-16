package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	t.Skip("リファクタリング中")

	// 0にすると適当な秋ポートが割り当てられる
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen %+v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return run(ctx)
	})

	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to get: %v", err)
	}

	expected := fmt.Sprintf("Hello, %s", in)
	if string(got) != expected {
		t.Errorf("want: %q, but got %q", expected, got)
	}

	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}

}
