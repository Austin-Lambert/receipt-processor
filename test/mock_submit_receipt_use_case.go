package test

import (
	"receipt-processor/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockSubmitReceiptUseCase struct {
	mock.Mock
}

func (m *MockSubmitReceiptUseCase) SubmitReceipt(receipt *domain.ReceiptDto) (uuid.UUID, error) {
	args := m.Called(receipt)
	return args.Get(0).(uuid.UUID), args.Error(1)
}
