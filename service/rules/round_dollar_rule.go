package rules

import "receipt-processor/domain"

type RoundDollarRule struct{}

func (r *RoundDollarRule) GetPoints(receipt *domain.Receipt) int {
	if receipt.Total%100 == 0 {
		return 50
	} else {
		return 0
	}
}
