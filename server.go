package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		l:   l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 引数のcontext使わなかったら、親のcancelが届かない
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
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
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Println("server shut down")
		return err
	}

	return eg.Wait()
}
