package service

import (
	"receipt-processor/domain"
	"receipt-processor/service/rules"
	"receipt-processor/storage"

	"github.com/google/uuid"
)

type ReceiptProcessorService struct {
	getReceiptProvider   storage.GetReceiptProvider
	storeReceiptProvider storage.StoreReceiptProvider
}

func NewReceiptProcessorService(getReceiptProvider storage.GetReceiptProvider, storeReceiptProvider storage.StoreReceiptProvider) (GetReceiptPointsUseCase, SubmitReceiptUseCase) {
	service := &ReceiptProcessorService{
		getReceiptProvider:   getReceiptProvider,
		storeReceiptProvider: storeReceiptProvider,
	}

	return service, service
}

func (s *ReceiptProcessorService) SubmitReceipt(receipt *domain.ReceiptDto) (uuid.UUID, error) {
	model, err := receipt.ToReceipt()
	if err != nil {
		return uuid.UUID{}, err
	}
	return s.storeReceiptProvider.StoreReceipt(model)
}

func (s *ReceiptProcessorService) GetReceiptPoints(id uuid.UUID) (int, error) {
	receipt, err := s.getReceiptProvider.GetReceipt(id)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, rule := range allRules {
		sum += rule.GetPoints(receipt)
	}
	return sum, nil
}

var allRules []rules.PointRule = []rules.PointRule{
	&rules.RetailerNameRule{},
	&rules.RoundDollarRule{},
	&rules.RoundQuarterRule{},
	&rules.ItemPairRule{},
	&rules.ItemDescNameRule{},
	&rules.OddPurchaseDayRule{},
	&rules.TwoToFourPurchaseTimeRule{},
}
