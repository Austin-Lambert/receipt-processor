package storage

import (
	"receipt-processor/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStoreAndGetReceipt(t *testing.T) {
	getRepo, storeRepo := NewReceiptRepository()
	receipt := &domain.Receipt{
		Retailer: "Costco",
	}
	receipt2 := &domain.Receipt{
		Retailer: "Target",
	}
	id, err := storeRepo.StoreReceipt(receipt)
	_, err2 := storeRepo.StoreReceipt(receipt2)
	assert.NoError(t, err)
	assert.NoError(t, err2)
	storedReceipt, err := getRepo.GetReceipt(id)
	assert.NoError(t, err)
	assert.NotNil(t, storedReceipt)
	assert.Equal(t, receipt, storedReceipt)
	assert.Equal(t, receipt.Retailer, storedReceipt.Retailer)
}

func TestGetReceiptNotFound(t *testing.T) {
	getRepo, _ := NewReceiptRepository()
	id := uuid.New()
	_, err := getRepo.GetReceipt(id)
	assert.Error(t, err)
	assert.Equal(t, err, ErrNotFound)
}
