package utils

type ArrayResponse struct {
	RecordCount int         `json:"record_count"`
	Records     interface{} `json:"record_list"`
}
