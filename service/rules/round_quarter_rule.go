package rules

import "receipt-processor/domain"

type RoundQuarterRule struct{}

func (r *RoundQuarterRule) GetPoints(receipt *domain.Receipt) int {
	if receipt.Total%25 == 0 {
		return 25
	} else {
		return 0
	}
}
