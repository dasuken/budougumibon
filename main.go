package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	if len(os.Args) != 2 {
		log.Print("need port number \n")
		os.Exit(1)
	}

	p := os.Args[1]
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", p))
	if err != nil {
		log.Fatalf("failed to listen port %s, err: %v \n", p, err)
	}
	if err := run(context.Background(), l); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

func run(ctx context.Context, l net.Listener) error {
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
		}),
	}

	// 引数のcontext使わなかったら、親のcancelが届かない
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}

		return nil
	})

	<-ctx.Done()
	// ここでctx渡さないのは何故なのかな
	// Shutdownで即キャンセル関数が走って`context canceledが帰ってくる`
	// shutdownはサーバー内部のコンテキストツリーにキャンセル処理を走らせhttp.ErrServerClosedを送信
	// 尚且つ引数のcontextにcancel処理
	// shutodownは内部実装的にctx.Doneで値を取得するとctx.Errを返し、なければ適当なtimeout(nextPoll)を設定してシャットダウン
	// nextPoll()のtickが具体的にどう決定されるのかはまだ見てない
	if err := s.Shutdown(context.Background()); err != nil {
		log.Println("server shut down")
		return err
	}

	return eg.Wait()
}
