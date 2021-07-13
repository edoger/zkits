package rest

import (
	"encoding/json"
)

type Response interface {
	Encode() ([]byte, error)
}

type StringResponse string

func (r StringResponse) Encode() ([]byte, error) {
	return []byte(r), nil
}

type JSONResponse struct {
	v interface{}
}

func (r *JSONResponse) Encode() ([]byte, error) {
	return json.Marshal(r.v)
}


