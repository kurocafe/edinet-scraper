# ビルドステージ
FROM golang:1.23-alpine AS builder

# 必要なパッケージのインストール
RUN apk add --no-cache git

WORKDIR /app

# 依存関係ファイルを先にコピー（キャッシュ効率化）
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# # ビルド（静的リンク）
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o edinet-scraper .

# # 実行ステージ（軽量イメージ）
# FROM alpine:latest

# # タイムゾーンとCA証明書のインストール
# RUN apk --no-cache add ca-certificates tzdata

# WORKDIR /root/

# # ビルドステージから実行ファイルをコピー
# COPY --from=builder /app/edinet-scraper .
# COPY --from=builder /app/.env.example .

# # 実行権限を付与
# RUN chmod +x ./edinet-scraper

# # デフォルトのエントリーポイント
# ENTRYPOINT ["./edinet-scraper"]

# # デフォルトの引数（上書き可能）
# CMD ["-date=2024-01-15"]
