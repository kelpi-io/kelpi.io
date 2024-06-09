package handlers

type Response struct {
	Result interface{} `json:"result"`
}

type DomainInfoResult struct {
	//ID             int32    `json:"id"`
	Zone string `json:"zone"`
	//Masters        []string `json:"masters"`
	//NotifiedSerial int32    `json:"notified_serial"`
	//Serial         uint32   `json:"serial"`
	//LastCheck      uint32   `json:"last_check"`
	//Kind           string   `json:"kind"`
}

type RecordInfo struct {
	QType   string `json:"qtype"`
	QName   string `json:"qname"`
	Content string `json:"content"`
	TTL     uint32 `json:"ttl"`
}

// type Backend interface {
// 	// Always required:
// 	Initialize() bool
// 	Lookup(qtype string, qname string, zoneId string) (lookupResultArray []LookupResult, err error)

// 	// Master operations:
// 	List(qname string, domainId string, zoneId string) (listResultArray []ListResult, err error)

// 	// Filling the Zone Cache:
// 	GetAllDomains(includeDisabled bool) (domainInfoResultArray []DomainInfoResult, err error)
// 	GetAllDomainMetadata(qname string) (metadata map[string][]string, err error)
// 	GetDomainMetadata(qname string, qtype string) (metadata []string, err error)
// 	GetDomainInfo(qname string) (domainInfo DomainInfoResult, err error)
// }
