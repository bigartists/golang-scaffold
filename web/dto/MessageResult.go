package dto

type MessageResult struct {
	Result  interface{} `jsn:"result"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}
