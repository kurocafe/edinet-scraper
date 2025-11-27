package config

import (
	"fmt"
	"os"
)

const (
	// APIEndpoint EDINET API v2のエンドポイント
	APIEndpoint = "https://api.edinet-fsa.go.jp/api/v2/documents.json"
	// APIVersion APIのバージョン
	APIVersion = "v2"
	// DefaultType 提出書類一覧及びメタデータ
	DefaultType = "2"
	// APIKeyEnvVar 環境変数名
	APIKeyEnvVar = "EDINET_API_KEY"
)

// GetAPIKey 環境変数からAPIキーを取得
func GetAPIKey() (string, error) {
	apiKey := os.Getenv(APIKeyEnvVar)
	if apiKey == "" {
		return "", fmt.Errorf("APIキーが設定されていません。環境変数 %s を設定してください", APIKeyEnvVar)
	}
	return apiKey, nil
}
