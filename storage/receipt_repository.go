package storage

import (
	"errors"
	"fmt"
	"receipt-processor/domain"

	"github.com/google/uuid"
)

type ReceiptRepository struct {
	receipts map[uuid.UUID]*domain.Receipt
}

func NewReceiptRepository() (GetReceiptProvider, StoreReceiptProvider) {
	repo := &ReceiptRepository{
		receipts: make(map[uuid.UUID]*domain.Receipt),
	}
	// I am aware this is a little weird, but I wanted to follow CQRS for this project.
	// If there was a DB here rather than a simple dictionary and/or if I was using a DI framework, this would make more sense.
	return repo, repo
}

var ErrNotFound = errors.New("No receipt found for that ID.")

func (r *ReceiptRepository) GetReceipt(id uuid.UUID) (*domain.Receipt, error) {
	receipt, exists := r.receipts[id]
	if !exists {
		fmt.Printf("Receipt with id %s not found\n", id)
		return nil, ErrNotFound
	}
	fmt.Printf("Receipt with id %s found\n", id)
	return receipt, nil
}

func (r *ReceiptRepository) StoreReceipt(receipt *domain.Receipt) (uuid.UUID, error) {
	id := uuid.New()
	r.receipts[id] = receipt
	fmt.Printf("Stored receipt with ID - %s\n", id)
	return id, nil
}
