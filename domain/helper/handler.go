package domain_helper

type ResponseError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
