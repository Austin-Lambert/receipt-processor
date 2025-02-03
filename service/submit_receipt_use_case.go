package service

import (
	"receipt-processor/domain"

	"github.com/google/uuid"
)

type SubmitReceiptUseCase interface {
	SubmitReceipt(receipt *domain.ReceiptDto) (uuid.UUID, error)
}
