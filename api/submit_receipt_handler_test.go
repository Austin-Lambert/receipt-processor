package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"receipt-processor/domain"
	"receipt-processor/storage"
	"receipt-processor/test"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubmitHappyPath(t *testing.T) {
	mockUseCase := new(test.MockSubmitReceiptUseCase)
	mockResponseWriter := new(test.MockResponseWriter)
	handler := NewSubmitReceiptHandler(mockUseCase, mockResponseWriter)

	testID := uuid.New()
	mockHttpResponseWriter := new(test.MockHttpResponseWriter)
	expected := SubmitResponse{testID.String()}
	input := getInputReceipt()
	mockUseCase.On("SubmitReceipt", input).Return(testID, nil)
	mockResponseWriter.On("WriteJSONResponse", mockHttpResponseWriter, http.StatusOK, expected).Return(nil)
	req, err := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewReader(getJsonInput(input)))
	if err != nil {
		t.Fatal(err)
	}
	result := handler.Handle(mockHttpResponseWriter, req)
	assert.NoError(t, result)
	mockUseCase.AssertCalled(t, "SubmitReceipt", input)
	mockResponseWriter.AssertCalled(t, "WriteJSONResponse", mockHttpResponseWriter, http.StatusOK, expected)
}

func TestSubmitWrongMethod(t *testing.T) {
	mockUseCase := new(test.MockSubmitReceiptUseCase)
	mockResponseWriter := new(test.MockResponseWriter)
	handler := NewSubmitReceiptHandler(mockUseCase, mockResponseWriter)

	mockHttpResponseWriter := new(test.MockHttpResponseWriter)
	req, err := http.NewRequest(http.MethodGet, "/receipts/process", nil)
	if err != nil {
		t.Fatal(err)
	}
	result := handler.Handle(mockHttpResponseWriter, req)
	assert.Error(t, result)
	mockUseCase.AssertNotCalled(t, "SubmitReceipt")
	mockResponseWriter.AssertNotCalled(t, "WriteJSONResponse", mockHttpResponseWriter, http.StatusOK, SubmitResponse{})
}

func TestSubmitParseError(t *testing.T) {
	mockUseCase := new(test.MockSubmitReceiptUseCase)
	mockResponseWriter := new(test.MockResponseWriter)
	handler := NewSubmitReceiptHandler(mockUseCase, mockResponseWriter)

	mockHttpResponseWriter := new(test.MockHttpResponseWriter)
	body := strings.NewReader(`{"retailer": "Store", "purchaseDate": "2022-01-01", "purchaseTime": 13:01`) // Missing closing quote and curly brace
	req, err := http.NewRequest(http.MethodPost, "/receipts/process", body)
	if err != nil {
		t.Fatal(err)
	}
	result := handler.Handle(mockHttpResponseWriter, req)
	assert.Error(t, result)
	mockUseCase.AssertNotCalled(t, "SubmitReceipt")
	mockResponseWriter.AssertNotCalled(t, "WriteJSONResponse", mockHttpResponseWriter, http.StatusOK, SubmitResponse{})
}

func TestSubmitReceiptError(t *testing.T) {
	mockUseCase := new(test.MockSubmitReceiptUseCase)
	mockResponseWriter := new(test.MockResponseWriter)
	handler := NewSubmitReceiptHandler(mockUseCase, mockResponseWriter)

	mockHttpResponseWriter := new(test.MockHttpResponseWriter)
	input := getInputReceipt()
	req, err := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewReader(getJsonInput(input)))
	if err != nil {
		t.Fatal(err)
	}

	mockUseCase.On("SubmitReceipt", mock.Anything).Return(uuid.Nil, storage.ErrNotFound)
	result := handler.Handle(mockHttpResponseWriter, req)
	assert.Error(t, result)
	mockUseCase.AssertCalled(t, "SubmitReceipt", input)
	mockResponseWriter.AssertNotCalled(t, "WriteJSONResponse", mockHttpResponseWriter, http.StatusOK, SubmitResponse{})
}

func getInputReceipt() *domain.ReceiptDto {
	return &domain.ReceiptDto{
		Retailer:     "Costco",
		PurchaseDate: "2020-01-01",
		PurchaseTime: "10:00",
		Items: []domain.ItemDto{
			{ShortDescription: "item1", Price: "10.00"},
		},
		Total: "10.00",
	}
}

func getJsonInput(input interface{}) []byte {
	jsonInput, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	return jsonInput
}
