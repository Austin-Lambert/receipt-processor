package storage

import (
	"receipt-processor/domain"

	"github.com/google/uuid"
)

type StoreReceiptProvider interface {
	StoreReceipt(receipt *domain.Receipt) (uuid.UUID, error)
}
