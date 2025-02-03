package api

import (
	"net/http"
	"receipt-processor/storage"
	"receipt-processor/test"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPointsHappyPath(t *testing.T) {
	mockUseCase := new(test.MockGetReceiptPointsUseCase)
	mockResponseWriter := new(test.MockResponseWriter)
	handler := NewGetReceiptPointsHandler(mockUseCase, mockResponseWriter)

	testID := uuid.New()
	expectedPoints := 100
	mockUseCase.On("GetReceiptPoints", testID).Return(expectedPoints, nil)
	mockResponseWriter.On("WriteJSONResponse", mock.Anything, http.StatusOK, mock.Anything).Return(nil)

	mockHttpResponseWriter := new(test.MockHttpResponseWriter)
	req, err := http.NewRequest(http.MethodGet, "/receipts/"+testID.String()+"/points", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("id", testID.String())
	result := handler.Handle(mockHttpResponseWriter, req)
	assert.NoError(t, result)
	mockUseCase.AssertCalled(t, "GetReceiptPoints", testID)
	mockResponseWriter.AssertCalled(t, "WriteJSONResponse", mockHttpResponseWriter, http.StatusOK, PointsResponse{expectedPoints})
}

func TestPointsWrongMethod(t *testing.T) {
	mockUseCase := new(test.MockGetReceiptPointsUseCase)
	mockResponseWriter := new(test.MockResponseWriter)
	handler := NewGetReceiptPointsHandler(mockUseCase, mockResponseWriter)

	mockHttpResponseWriter := new(test.MockHttpResponseWriter)
	req, err := http.NewRequest(http.MethodPost, "/receipts/"+uuid.New().String()+"/points", nil)
	if err != nil {
		t.Fatal(err)
	}

	result := handler.Handle(mockHttpResponseWriter, req)
	assert.Error(t, result)
	mockUseCase.AssertNotCalled(t, "GetReceiptPoints")
	mockResponseWriter.AssertNotCalled(t, "WriteJSONResponse", mockHttpResponseWriter, http.StatusOK, PointsResponse{})
}

func TestPointsBadUUID(t *testing.T) {
	mockUseCase := new(test.MockGetReceiptPointsUseCase)
	mockResponseWriter := new(test.MockResponseWriter)
	handler := NewGetReceiptPointsHandler(mockUseCase, mockResponseWriter)

	mockHttpResponseWriter := new(test.MockHttpResponseWriter)
	req, err := http.NewRequest(http.MethodGet, "/receipts/bad-uuid/points", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("id", "bad-uuid")

	result := handler.Handle(mockHttpResponseWriter, req)
	assert.Error(t, result)
	mockUseCase.AssertNotCalled(t, "GetReceiptPoints")
	mockResponseWriter.AssertNotCalled(t, "WriteJSONResponse", mockHttpResponseWriter, http.StatusOK, PointsResponse{})
}

func TestPointsReceiptNotFound(t *testing.T) {
	mockUseCase := new(test.MockGetReceiptPointsUseCase)
	mockResponseWriter := new(test.MockResponseWriter)
	handler := NewGetReceiptPointsHandler(mockUseCase, mockResponseWriter)
	testID := uuid.New()
	mockUseCase.On("GetReceiptPoints", testID).Return(0, storage.ErrNotFound)

	mockHttpResponseWriter := new(test.MockHttpResponseWriter)
	req, err := http.NewRequest(http.MethodGet, "/receipts/"+uuid.New().String()+"/points", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("id", testID.String())

	result := handler.Handle(mockHttpResponseWriter, req)
	assert.Error(t, result)
	mockUseCase.AssertCalled(t, "GetReceiptPoints", mock.Anything)
	mockResponseWriter.AssertNotCalled(t, "WriteJSONResponse", mockHttpResponseWriter, http.StatusOK, PointsResponse{})
}
