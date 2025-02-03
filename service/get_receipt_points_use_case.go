package service

import "github.com/google/uuid"

type GetReceiptPointsUseCase interface {
	GetReceiptPoints(id uuid.UUID) (int, error)
}
