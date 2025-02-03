package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockResponseWriter struct {
	mock.Mock
}

func (m *MockResponseWriter) WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	args := m.Called(w, statusCode, data)
	return args.Error(0)
}
