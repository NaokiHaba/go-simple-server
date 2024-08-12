package main

import (
	"fmt"
	"log"

	"github.com/NaokiHaba/go-simple-server/server"
)

const (
	port = 8080
	host = "0.0.0.0"
)

func main() {
	// サーバーの新しいインスタンスを作成
	s := server.New()

	// サーバーを起動
	// ホストとポートを組み合わせてアドレスを作成
	err := s.Start(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		// エラーが発生した場合、サーバーを停止
		s.Stop()
		// エラーメッセージをログに記録して終了
		log.Fatalf("Failed to start server: %v", err)
	}
}
