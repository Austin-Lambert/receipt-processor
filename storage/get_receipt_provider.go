package storage

import (
	"receipt-processor/domain"

	"github.com/google/uuid"
)

type GetReceiptProvider interface {
	GetReceipt(id uuid.UUID) (*domain.Receipt, error)
}
