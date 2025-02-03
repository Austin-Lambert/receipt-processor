package rules

import "receipt-processor/domain"

type ItemPairRule struct{}

func (r *ItemPairRule) GetPoints(receipt *domain.Receipt) int {
	items := len(receipt.Items)
	pairs := items / 2
	return pairs * 5
}
