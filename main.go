package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kurocafe/edinet-scraper/api"
	"github.com/kurocafe/edinet-scraper/config"
	"github.com/kurocafe/edinet-scraper/models"
)

func main() {
	// .envファイルの読み込み（存在しない場合はスキップ）
	if err := godotenv.Load(); err != nil {
		log.Println(".envファイルが見つかりません。環境変数から読み込みます。")
	}

	// コマンドライン引数の定義
	date := flag.String("date", "", "取得する日付 (YYYY-MM-DD形式)")
	flag.Parse()

	// 日付の検証
	if *date == "" {
		fmt.Println("使い方: go run main.go -date=YYYY-MM-DD")
		fmt.Println("例: go run main.go -date=2024-01-15")
		os.Exit(1)
	}

	// APIキーの取得
	apiKey, err := config.GetAPIKey()
	if err != nil {
		log.Fatalf("エラー: %v", err)
	}

	// 書類一覧の取得
	fmt.Printf("日付 %s の書類一覧を取得中...\n\n", *date)
	response, err := api.FetchDocumentList(*date, apiKey, config.APIEndpoint, config.DefaultType)
	if err != nil {
		log.Fatalf("書類一覧の取得に失敗しました: %v", err)
	}

	// 結果の表示
	displayResults(response)
}

// displayResults 取得した書類一覧を見やすく表示
func displayResults(response *models.DocumentListResponse) {
	fmt.Println("========================================")
	fmt.Printf("取得件数: %d件\n", response.Metadata.ResultSet.Count)
	fmt.Printf("処理日時: %s\n", response.Metadata.ProcessTime)
	fmt.Println("========================================\n")

	if len(response.Results) == 0 {
		fmt.Println("該当する書類がありませんでした。")
		return
	}

	// 各書類情報の表示
	for i, result := range response.Results {
		fmt.Printf("[%d] %s\n", i+1, result.FilerName)
		fmt.Printf("    証券コード: %s\n", result.SecCode)
		fmt.Printf("    EDINETコード: %s\n", result.EDINETCode)
		fmt.Printf("    書類種別: %s\n", result.DocDescription)
		fmt.Printf("    提出日時: %s\n", result.SubmitDateTime)
		fmt.Printf("    書類ID: %s\n", result.DocID)
		fmt.Println()
	}
}
