package resolvers

type RecordInfo struct {
	QType   string `json:"qtype"`
	QName   string `json:"qname"`
	Content string `json:"content"`
	TTL     int    `json:"ttl"`
}

type Response struct {
	Result interface{} `json:"result"`
}

type DomainInfoResult struct {
	Zone string `json:"zone"`
}
