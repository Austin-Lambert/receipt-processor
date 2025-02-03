package service

import (
	"fmt"
	"receipt-processor/domain"
	"receipt-processor/test"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubmitReceipt(t *testing.T) {
	mockStoreReceiptProvider := new(test.MockStoreReceiptProvider)
	mockGetReceiptProvider := new(test.MockGetReceiptProvider)
	_, storeService := NewReceiptProcessorService(mockGetReceiptProvider, mockStoreReceiptProvider)

	input := getTargetTestReceipt().ToReceiptDto()
	fmt.Printf("Input: %+v\n", input)
	expected, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	mockStoreReceiptProvider.On("StoreReceipt", mock.Anything).Return(expected, nil)
	result, err := storeService.SubmitReceipt(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetPoints(t *testing.T) {
	mockStoreReceiptProvider := new(test.MockStoreReceiptProvider)
	mockGetReceiptProvider := new(test.MockGetReceiptProvider)
	getService, _ := NewReceiptProcessorService(mockGetReceiptProvider, mockStoreReceiptProvider)
	input, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	mockGetReceiptProvider.On("GetReceipt", input).Return(getTargetTestReceipt(), nil)
	points, err := getService.GetReceiptPoints(input)
	assert.NoError(t, err)
	assert.Equal(t, 28, points)
}

func getTargetTestReceipt() *domain.Receipt {
	items := []domain.Item{
		{
			ShortDescription: "Mountain Dew 12PK",
			Price:            649,
		},
		{
			ShortDescription: "Emils Cheese Pizza",
			Price:            1225,
		},
		{
			ShortDescription: "Knorr Creamy Chicken",
			Price:            126,
		},
		{
			ShortDescription: "Doritos Nacho Cheese",
			Price:            335,
		},
		{
			ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
			Price:            1200,
		},
	}

	timestamp, err := time.Parse("2006-01-02 15:04", "2022-01-01 13:01")
	if err != nil {
		panic(err)
	}

	return &domain.Receipt{
		Retailer:          "Target",
		PurchaseTimestamp: timestamp,
		Items:             items,
		Total:             3535,
	}
}
