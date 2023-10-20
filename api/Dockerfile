# ベースイメージを指定
FROM golang:1.21.0 as build

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# ソースコードをコンテナ内にコピー
COPY . .

# Goアプリケーションのビルド
RUN go build -o main .

# ポート8000を公開
EXPOSE 8000

# アプリケーションを起動
CMD ["./main"]
