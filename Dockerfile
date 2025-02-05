FROM golang:1.23

# モジュールを使用して依存関係を管理
ENV GO111MODULE=on

# ワーキングディレクトリの設定
WORKDIR /app

# ソースコードとairの設定ファイルのコピー
COPY . ./

# airのインストール
RUN go mod tidy
RUN go install github.com/air-verse/air@latest

# ポートを開放
EXPOSE 1323

# airでアプリケーションを起動
CMD ["/go/bin/air"]

