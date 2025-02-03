package test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockHttpResponseWriter struct {
	mock.Mock
}

func (m *MockHttpResponseWriter) Header() http.Header {
	args := m.Called()
	return args.Get(0).(http.Header)
}

func (m *MockHttpResponseWriter) Write(p []byte) (int, error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *MockHttpResponseWriter) WriteHeader(statusCode int) {
	m.Called(statusCode)
}
