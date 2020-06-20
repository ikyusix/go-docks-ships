package model

type ResponseWrapper struct {
	Success bool
	Message string
	Data interface{} `json:"data,omitempty"`
}