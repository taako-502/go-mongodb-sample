build:
	go build -o app/bin/app app/main.go

# デフォルトターゲットを設定
.PHONY: build