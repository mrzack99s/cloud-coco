package types

type ArrayResponse struct {
	RecordCount int64       `json:"record_count"`
	Records     interface{} `json:"record_list"`
}
