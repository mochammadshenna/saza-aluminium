package api

import "github.com/mochammadshenna/saza-aluminium/util/helper/json"



type WebResponse struct {
	Code   int
	Status string
	Data   interface{}
}

type ApiResponse struct {
	Header HeaderResponse `json:"header"`
	Data   interface{}    `json:"data"`
	Error  interface{}    `json:"error"`
}

func (a *ApiResponse) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}

type HeaderResponse struct {
	ServerTimeMs  int64  `json:"serverTimeMs"`
	ProcessTimeMs int64  `json:"processTimeMs"`
	RequestId     string `json:"requestId"`
}

type HttpError interface {
	Error() string
	StatusCode() int
}

type ErrorResponse struct {
	HttpCode int         `json:"-"`
	Code     string      `json:"code"`
	Message  interface{} `json:"message"`
	Errors   interface{} `json:"errors"`
}

func (e ErrorResponse) Error() string {
	return e.Message.(string)
}

func (e ErrorResponse) StatusCode() int {
	return e.HttpCode
}

type ErrorValidate struct {
	Key     string `json:"key"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
