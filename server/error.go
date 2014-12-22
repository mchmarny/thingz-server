package server

type JSONError struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
