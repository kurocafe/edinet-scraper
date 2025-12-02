package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/kurocafe/edinet-scraper/models"
)

// FetchDocumentList 指定した日付の書類一覧を取得
func FetchDocumentList(date string, apiKey string, endPoint string, tp string) (*models.DocumentListResponse, error) {
	// URLパラメータの構築
	params := url.Values{}
	params.Add("date", date)
	params.Add("type", tp)
	params.Add("Subscription-Key", apiKey)

	// リクエストURLの構築
	requestURL := fmt.Sprintf("%s?%s", endPoint, params.Encode())

	// HTTPリクエストの作成
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("リクエストの作成に失敗しました: %w", err)
	}

	// ヘッダーにAPIキーを設定
	req.Header.Set("Subscription-Key", apiKey)

	// HTTPクライアントの作成とリクエストの送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("APIリクエストに失敗しました: %w", err)
	}
	defer resp.Body.Close()

	// ステータスコードのチェック
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("APIエラー (ステータスコード: %d): %s", resp.StatusCode, string(body))
	}

	// レスポンスボディの読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("レスポンスの読み取りに失敗しました: %w", err)
	}

	fmt.Println(body)
	var result map[string]interface{}
	if err_test := json.Unmarshal(body, &result); err_test != nil {
		return nil, fmt.Errorf("JSONのパースに失敗しました: %w", err_test)
	}

	fmt.Printf("%+v\n", result)

	// JSONのパース
	var response models.DocumentListResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("JSONのパースに失敗しました: %w", err)
	}

	fmt.Printf("%+v\n", response)

	// APIステータスのチェック
	if response.Metadata.Status != "200" {
		return nil, fmt.Errorf("APIエラー: %s", response.Metadata.Message)
	}

	return &response, nil
}
