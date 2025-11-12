package router

import (
	"encoding/json/v2"
	"net/http"
)

type response[T any] struct {
	Success bool   `json:"success"`
	Data    *T     `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func writeResponse[T any](w http.ResponseWriter, r *response[T], status int, opts ...json.Options) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.MarshalWrite(w, *r, opts...)
}

func writeSuccess[T any](w http.ResponseWriter, data T, status int, opts ...json.Options) error {
	return writeResponse(w, &response[T]{Success: true, Data: &data}, status, opts...)
}

func writeError(w http.ResponseWriter, err error, status int, opts ...json.Options) error {
	return writeResponse(w, &response[any]{Success: false, Error: err.Error()}, status, opts...)
}
