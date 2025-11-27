# 設計書：EDINET APIスクレイパー

## 1. システム構成

```
┌─────────────┐
│   ユーザー   │
└──────┬──────┘
       │
       ▼
┌─────────────────────┐
│  main.go            │
│  (エントリーポイント)  │
└──────┬──────────────┘
       │
       ▼
┌─────────────────────┐
│  APIクライアント      │
│  - リクエスト送信     │
│  - レスポンス受信     │
└──────┬──────────────┘
       │
       ▼
┌─────────────────────┐
│  EDINET API (v2)    │
│  (金融庁提供)        │
└─────────────────────┘
```

## 2. 使用する主要パッケージ

- `net/http`: HTTPリクエストの送信
- `encoding/json`: JSONのパース
- `time`: 日付の処理
- `fmt`: 出力
- `sync`: 並行処理の制御（WaitGroupなど）

## 3. データ構造

### 3.1 APIレスポンスの構造体

```go
// 書類一覧APIのレスポンス全体
type DocumentListResponse struct {
    Metadata Metadata  `json:"metadata"`
    Results  []Result  `json:"results"`
}

// メタデータ
type Metadata struct {
    Title       string      `json:"title"`
    Parameter   Parameter   `json:"parameter"`
    ResultSet   ResultSet   `json:"resultset"`
    ProcessTime string      `json:"processDateTime"`
    Status      string      `json:"status"`
    Message     string      `json:"message"`
}

// リクエストパラメータ
type Parameter struct {
    Date string `json:"date"`
    Type string `json:"type"`
}

// 結果セット
type ResultSet struct {
    Count int `json:"count"`
}

// 個別の書類情報
type Result struct {
    SeqNumber        int    `json:"seqNumber"`
    DocID            string `json:"docID"`
    EDINETCode       string `json:"edinetCode"`
    SecCode          string `json:"secCode"`
    JCN              string `json:"JCN"`
    FilerName        string `json:"filerName"`
    FundCode         string `json:"fundCode"`
    OrdinanceCode    string `json:"ordinanceCode"`
    FormCode         string `json:"formCode"`
    DocTypeCode      string `json:"docTypeCode"`
    PeriodStart      string `json:"periodStart"`
    PeriodEnd        string `json:"periodEnd"`
    SubmitDateTime   string `json:"submitDateTime"`
    DocDescription   string `json:"docDescription"`
    IssuerEDINETCode string `json:"issuerEdinetCode"`
    SubjectEDINETCode string `json:"subjectEdinetCode"`
    Subsidiary       string `json:"subsidiary"`
    CurrentReportReason string `json:"currentReportReason"`
    ParentDocID      string `json:"parentDocID"`
    OpeDateTime      string `json:"opeDateTime"`
    WithdrawalStatus string `json:"withdrawalStatus"`
    DocInfoEditStatus string `json:"docInfoEditStatus"`
    DisclosureStatus string `json:"disclosureStatus"`
    XBRLFlag         string `json:"xbrlFlag"`
    PDFFlag          string `json:"pdfFlag"`
    AttachDocFlag    string `json:"attachDocFlag"`
    EnglishDocFlag   string `json:"englishDocFlag"`
}
```

## 4. 主要な関数設計

### 4.1 main関数
```go
func main()
```
**役割**: プログラムのエントリーポイント
**処理**:
1. APIキーの読み込み（環境変数または設定ファイル）
2. 取得対象の日付リストを準備
3. fetchDocuments関数を呼び出し
4. 結果の表示

### 4.2 API呼び出し関数
```go
func fetchDocumentList(date string, apiKey string) (*DocumentListResponse, error)
```
**役割**: 指定した日付の書類一覧を取得
**引数**:
- `date`: 取得する日付（YYYY-MM-DD形式）
- `apiKey`: EDINET APIキー
**戻り値**:
- `*DocumentListResponse`: パースされたレスポンス
- `error`: エラー情報

**処理フロー**:
1. URLとパラメータの構築
2. HTTPリクエストの送信
3. レスポンスの取得
4. JSONのパース
5. エラーチェック

### 4.3 並行処理関数（Phase 2で実装）
```go
func fetchMultipleDates(dates []string, apiKey string) ([]DocumentListResponse, error)
```
**役割**: 複数の日付のデータを並行処理で取得
**引数**:
- `dates`: 日付のスライス
- `apiKey`: EDINET APIキー
**戻り値**:
- `[]DocumentListResponse`: 取得した結果のスライス
- `error`: エラー情報

**並行処理の方針**:
- goroutineで各日付のリクエストを並行実行
- channelで結果を受け取る
- sync.WaitGroupで全てのgoroutineの完了を待つ
- APIレート制限を考慮してsleepを入れる

### 4.4 表示関数
```go
func displayResults(response *DocumentListResponse)
```
**役割**: 取得した書類一覧を見やすく表示
**引数**:
- `response`: 取得したレスポンス

**表示項目**:
- 提出者名
- 証券コード
- 書類種別
- 提出日時

## 5. エラーハンドリング方針

### 5.1 エラーの種類と対応

| エラーの種類 | 対応方法 |
|-------------|---------|
| ネットワークエラー | エラーメッセージを表示し、処理を終了 |
| APIキー未設定 | 明確なエラーメッセージを表示 |
| JSONパースエラー | レスポンスの内容をログに出力し、処理を終了 |
| HTTPステータスエラー | ステータスコードとメッセージを表示 |

### 5.2 エラーメッセージの方針
- ユーザーが理解しやすい日本語メッセージ
- 問題の原因と対処方法を含める
- デバッグ用の詳細情報も出力

## 6. 設定管理

### 6.1 定数
```go
const (
    APIEndpoint = "https://api.edinet-fsa.go.jp/api/v2/documents.json"
    APIVersion  = "v2"
    DefaultType = "2" // 提出書類一覧及びメタデータ
)
```

### 6.2 APIキーの管理方法
- 環境変数 `EDINET_API_KEY` から読み込む
- または、コマンドライン引数で渡す

## 7. テスト方針

### 7.1 単体テストの対象
- JSONパース機能
- URL構築機能
- 日付フォーマット変換

### 7.2 統合テストの方針
- 実際のAPIを使ったエンドツーエンドテスト
- テスト用の固定日付を使用

## 8. 今後の拡張ポイント

### 8.1 フィルタリング機能
- 書類種別でフィルタ（有価証券報告書のみ、など）
- 企業コードでフィルタ

### 8.2 出力機能
- CSV出力
- JSON出力
- データベース保存

### 8.3 書類取得機能
- 書類取得APIの実装
- ZIP/PDFファイルのダウンロード

## 9. ディレクトリ構成

```
edinet-scraper/
├── main.go              # エントリーポイント
├── api/
│   └── client.go       # APIクライアント実装
├── models/
│   └── document.go     # データ構造定義
├── config/
│   └── config.go       # 設定管理
├── go.mod              # Go modules
└── README.md           # 使い方の説明
```

## 10. 実装の優先順位

1. **最優先**: 単一日付の書類一覧取得（Phase 1）
2. **次**: 並行処理の実装（Phase 2）
3. **その後**: フィルタリングと出力機能（Phase 3, 4）
