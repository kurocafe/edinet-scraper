# EDINET APIスクレイパー

EDINET APIを使用して、企業の有価証券報告書などの書類情報を取得するツールです。

## 機能

- 指定した日付の提出書類一覧を取得
- 企業名、証券コード、書類種別などの基本情報を表示

## セットアップ

### 1. EDINET APIキーの取得

[EDINET API利用登録ページ](https://disclosure2.edinet-fsa.go.jp/)でAPIキーを取得してください。

### 2. 環境変数の設定

取得したAPIキーを設定します。以下の2つの方法があります：

#### 方法1：.envファイルを使用（推奨）

```bash
# .env.exampleをコピーして.envファイルを作成
cp .env.example .env

# .envファイルを編集してAPIキーを設定
# EDINET_API_KEY=your-api-key-here
```

#### 方法2：環境変数に直接設定

```bash
export EDINET_API_KEY="your-api-key-here"
```

## 使い方

### 方法1：Dockerを使う（推奨）

Dockerを使えば、Goのインストール不要で実行できます。

```bash
# Dockerイメージをビルド
docker build -t edinet-scraper .

# 実行（.envファイルから環境変数を読み込む）
docker run --rm --env-file .env edinet-scraper -date=2024-01-15

# または docker-compose を使う
docker-compose run --rm edinet-scraper -date=2024-01-15
```

### 方法2：ローカルで実行

```bash
# 直接実行
go run main.go -date=2024-01-15

# またはビルドしてから実行
go build -o edinet-scraper
./edinet-scraper -date=2024-01-15
```

### 実行例

```bash
# 2024年1月15日の書類一覧を取得
docker run --rm --env-file .env edinet-scraper -date=2024-01-15

# 出力例：
# 日付 2024-01-15 の書類一覧を取得中...
#
# ========================================
# 取得件数: 50件
# 処理日時: 2024-01-15 18:00:00
# ========================================
#
# [1] トヨタ自動車株式会社
#     証券コード: 7203
#     EDINETコード: E01234
#     書類種別: 有価証券報告書
#     提出日時: 2024-01-15 17:00:00
#     書類ID: S100ABCD
```

## プロジェクト構成

```
edinet-scraper/
├── main.go              # エントリーポイント
├── api/
│   └── client.go       # APIクライアント実装
├── models/
│   └── document.go     # データ構造定義
├── config/
│   └── config.go       # 設定管理
├── Dockerfile           # Docker設定
├── docker-compose.yml   # Docker Compose設定
├── .dockerignore        # Dockerビルド除外ファイル
├── .env.example         # 環境変数テンプレート
├── go.mod              # Go modules
├── go.sum              # Go modules checksum
├── design.md           # 設計書
├── requirements.md     # 要件定義書
└── README.md           # このファイル
```

## 開発ロードマップ

- [x] **Phase 1**: 単一日付の書類一覧取得（基本機能）
- [ ] **Phase 2**: 複数日付の並行処理対応
- [ ] **Phase 3**: フィルタリング機能の追加
- [ ] **Phase 4**: ファイル出力機能の追加

## ライセンス

MIT