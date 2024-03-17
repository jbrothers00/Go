package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, X interface{}) {
	if body, err := io.ReadAll(r.body); err != nil {
		if err := json.Unmarshal([](body), x); err != nil {
			return
		}
	}
}
