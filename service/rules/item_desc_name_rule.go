package rules

import (
	"math"
	"receipt-processor/domain"
	"strings"
)

type ItemDescNameRule struct{}

func (r *ItemDescNameRule) GetPoints(receipt *domain.Receipt) int {
	sum := 0
	for _, item := range receipt.Items {
		trimmed := strings.TrimSpace(item.ShortDescription)
		trimmedLen := len(trimmed)
		if trimmedLen%3 == 0 {
			points := math.Ceil(float64(item.Price) * 0.2 / 100)
			sum += int(points)
		}
	}
	return sum
}
