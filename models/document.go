package models

// DocumentListResponse 書類一覧APIのレスポンス全体
type DocumentListResponse struct {
	Metadata Metadata `json:"metadata"`
	Results  []Result `json:"results"`
}

// Metadata メタデータ
type Metadata struct {
	Title       string    `json:"title"`
	Parameter   Parameter `json:"parameter"`
	ResultSet   ResultSet `json:"resultset"`
	ProcessTime string    `json:"processDateTime"`
	Status      string    `json:"status"`
	Message     string    `json:"message"`
}

// Parameter リクエストパラメータ
type Parameter struct {
	Date string `json:"date"`
	Type string `json:"type"`
}

// ResultSet 結果セット
type ResultSet struct {
	Count int `json:"count"`
}

// Result 個別の書類情報
type Result struct {
	SeqNumber           int    `json:"seqNumber"`
	DocID               string `json:"docID"`
	EDINETCode          string `json:"edinetCode"`
	SecCode             string `json:"secCode"`
	JCN                 string `json:"JCN"`
	FilerName           string `json:"filerName"`
	FundCode            string `json:"fundCode"`
	OrdinanceCode       string `json:"ordinanceCode"`
	FormCode            string `json:"formCode"`
	DocTypeCode         string `json:"docTypeCode"`
	PeriodStart         string `json:"periodStart"`
	PeriodEnd           string `json:"periodEnd"`
	SubmitDateTime      string `json:"submitDateTime"`
	DocDescription      string `json:"docDescription"`
	IssuerEDINETCode    string `json:"issuerEdinetCode"`
	SubjectEDINETCode   string `json:"subjectEdinetCode"`
	Subsidiary          string `json:"subsidiary"`
	CurrentReportReason string `json:"currentReportReason"`
	ParentDocID         string `json:"parentDocID"`
	OpeDateTime         string `json:"opeDateTime"`
	WithdrawalStatus    string `json:"withdrawalStatus"`
	DocInfoEditStatus   string `json:"docInfoEditStatus"`
	DisclosureStatus    string `json:"disclosureStatus"`
	XBRLFlag            string `json:"xbrlFlag"`
	PDFFlag             string `json:"pdfFlag"`
	AttachDocFlag       string `json:"attachDocFlag"`
	EnglishDocFlag      string `json:"englishDocFlag"`
}
