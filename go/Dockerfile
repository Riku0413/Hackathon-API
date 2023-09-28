# ベースイメージを指定
FROM golang:1.21 as build

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# ソースコードをコンテナ内にコピー
COPY main.go .
COPY go.mod .

# Goアプリケーションのビルド
RUN go build -o main .

# ポート8000を公開
EXPOSE 8080

# アプリケーションを起動
CMD ["./main"]