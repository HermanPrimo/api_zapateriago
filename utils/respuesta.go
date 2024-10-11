package utils

type Respuesta struct {
	Msg        string      `json:"message"`
	StatusCode int         `json:"code"`
	Data       interface{} `json:"data"`
}
