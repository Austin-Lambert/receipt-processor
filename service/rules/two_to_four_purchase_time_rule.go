package rules

import "receipt-processor/domain"

type TwoToFourPurchaseTimeRule struct{}

func (r *TwoToFourPurchaseTimeRule) GetPoints(receipt *domain.Receipt) int {
	if receipt.PurchaseTimestamp.Hour() >= 14 && receipt.PurchaseTimestamp.Hour() <= 16 {
		return 10
	} else {
		return 0
	}
}
