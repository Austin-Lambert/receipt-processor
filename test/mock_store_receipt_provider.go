package test

import (
	"receipt-processor/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockStoreReceiptProvider struct {
	mock.Mock
}

func (m *MockStoreReceiptProvider) StoreReceipt(receipt *domain.Receipt) (uuid.UUID, error) {
	args := m.Called(receipt)
	return args.Get(0).(uuid.UUID), args.Error(1)
}
