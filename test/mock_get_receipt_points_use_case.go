package test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockGetReceiptPointsUseCase struct {
	mock.Mock
}

func (m *MockGetReceiptPointsUseCase) GetReceiptPoints(id uuid.UUID) (int, error) {
	args := m.Called(id)
	return args.Get(0).(int), args.Error(1)
}
