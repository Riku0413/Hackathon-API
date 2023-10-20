# ベースイメージを指定
FROM golang:1.21.0 as build

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# Go モジュールの依存関係をコピー
COPY go.mod .
COPY go.sum .

# Go モジュールをインストール
RUN go mod download

# ソースコードをコンテナ内にコピー
COPY . .

# Go アプリケーションのビルド
RUN go build -o main .

# ポート 8000 を公開
EXPOSE 8000

# アプリケーションを起動
CMD ["./main"]