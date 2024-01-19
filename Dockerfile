FROM golang:1.21.3

# モジュールを使用して依存関係を管理
ENV GO111MODULE=on

# ワーキングディレクトリの設定
WORKDIR /app

# ソースコードとairの設定ファイルのコピー
COPY . ./

# airのインストール
RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest

# ポートを開放
EXPOSE 1323

# airでアプリケーションを起動
CMD ["air"]
