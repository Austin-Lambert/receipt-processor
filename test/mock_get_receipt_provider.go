package test

import (
	"receipt-processor/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockGetReceiptProvider struct {
	mock.Mock
}

func (m *MockGetReceiptProvider) GetReceipt(id uuid.UUID) (*domain.Receipt, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Receipt), args.Error(1)
}
