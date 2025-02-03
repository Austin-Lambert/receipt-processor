package api

import (
	"encoding/json"
	"net/http"
)

type ResponseWriter interface {
	WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) error
}

type DefaultResponseWriter struct {
}

// WriteJSONResponse implements ResponseWriter.
func (r DefaultResponseWriter) WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}
