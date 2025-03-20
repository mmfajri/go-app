package responses

type BaseResponse struct {
	StatusCode int `json:"status_code"`
	Message	string `json:"message"`
	Data any `json:"Data"`
}
