package jango

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string
	Msg    string
	Data   any
}

func StandardJsonResponse(writer http.ResponseWriter, status, msg string, data any, status_code ...int) error {
	response := &Response{
		Status: status,
		Msg:    msg,
		Data:   data,
	}
	return JsonResponse(writer, response, status_code...)
}

func JsonResponse(writer http.ResponseWriter, data any, status_codes ...int) error {
	status_code := 200
	if len(status_codes) > 0 {
		status_code = status_codes[0]
	}
	writer.WriteHeader(status_code)
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		return err
	}
	return nil
}
