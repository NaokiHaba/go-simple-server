# ビルドステージ：Go言語の開発環境を使用
FROM golang:1.19-alpine3.17 AS builder
# 作業ディレクトリを/appに設定し、以降の操作の基準点とする
WORKDIR /app
# カレントディレクトリのすべてのファイルをコンテナの作業ディレクトリにコピー
COPY . .
# go.modファイルに記載された依存関係をダウンロードし、キャッシュ
RUN go mod download
# main.goファイルをコンパイルし、実行可能な'main'バイナリを生成
RUN go build -o main /app/main.go

# 実行ステージ：軽量なAlpineイメージを使用し、最終的なイメージサイズを削減
FROM alpine:3.17
# 実行ステージでの作業ディレクトリを/appに設定
WORKDIR /app
# ビルドステージで生成した'main'バイナリを実行ステージの/appディレクトリにコピー
COPY --from=builder /app/main .
# コンテナがリッスンするポート8080を明示的に宣言
EXPOSE 8080
# コンテナ起動時に実行されるコマンドを指定（アプリケーションの起動）
CMD ["/app/main"]