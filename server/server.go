package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Server はサーバー構造体です
// counter: リクエスト数をカウントする変数
// server: HTTPサーバーのインスタンス
// router: Ginルーターのインスタンス
type Server struct {
	counter int64
	server  *http.Server
	router  *gin.Engine
}

// New は新しいサーバーを作成します
// 1. Ginのデフォルトルーターを初期化
// 2. Server構造体を初期化（カウンターを0に設定）
// 3. ルートとヘルスチェックのエンドポイントを設定
// 4. 初期化されたServerのポインタを返す
func New() *Server {
	router := gin.Default()
	server := &Server{
		counter: int64(0),
		router:  router,
	}

	router.GET("/", server.CounterHandler)
	router.GET("/health_checks", server.HealthHandler)

	return server
}

// HealthHandler はヘルスチェックエンドポイントのハンドラーです
// 200 OKレスポンスを返します
func (s *Server) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// CounterHandler はカウンターを増加させ、その値を返すハンドラーです
// 1. カウンターをインクリメント
// 2. 現在のカウンター値をJSONで返す
func (s *Server) CounterHandler(c *gin.Context) {
	s.counter++
	c.JSON(http.StatusOK, gin.H{
		"counter": s.counter,
	})
}

// Start はサーバーを起動します
// 1. HTTPサーバーを設定（アドレス、ハンドラー、読み取りタイムアウト）
// 2. サーバー起動のログを出力
// 3. サーバーを起動し、エラーがあれば返す
func (s *Server) Start(address string) error {
	s.server = &http.Server{
		Addr:        address,
		Handler:     s.router,
		ReadTimeout: 10 * time.Second,
	}

	log.Printf("Server is running on %s\n", address)

	return s.server.ListenAndServe()
}

// Stop はサーバーを停止します
// 1. 停止処理開始のログを出力
// 2. サーバーが未初期化の場合は何もせずに終了
// 3. サーバーを閉じる
func (s *Server) Stop() error {
	log.Println("Stopping server...")

	if s.server == nil {
		return nil
	}

	return s.server.Close()
}
