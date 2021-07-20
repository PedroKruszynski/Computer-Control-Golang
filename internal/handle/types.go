package handle

type Response struct {
	HttpCode int    `json:"httpCode"`
	Message  string `json:"message"`
}
