package rules

import (
	"receipt-processor/domain"
	"unicode"
)

type RetailerNameRule struct{}

func (r *RetailerNameRule) GetPoints(receipt *domain.Receipt) int {
	return countAlphanumeric(receipt.Retailer)
}

func countAlphanumeric(s string) int {
	count := 0
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			count++
		}
	}
	return count
}
