package rules

import "receipt-processor/domain"

type OddPurchaseDayRule struct{}

func (r *OddPurchaseDayRule) GetPoints(receipt *domain.Receipt) int {
	if receipt.PurchaseTimestamp.Day()%2 == 1 {
		return 6
	} else {
		return 0
	}
}
