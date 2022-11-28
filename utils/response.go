package utils

// Result generic structure for API json responses.
type Result struct {
	Code int `json:"code"`
	Data any `json:"data"`
}
