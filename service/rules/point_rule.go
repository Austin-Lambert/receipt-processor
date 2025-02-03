package rules

import "receipt-processor/domain"

type PointRule interface {
	GetPoints(receipt *domain.Receipt) int
}
