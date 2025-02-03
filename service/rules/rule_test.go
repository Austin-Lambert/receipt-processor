package rules

import (
	"receipt-processor/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRetailerNameRule(t *testing.T) {
	r := &RetailerNameRule{}
	receipt := getTargetTestReceipt()
	assert.Equal(t, 6, r.GetPoints(receipt))

	receipt.Retailer = "Walmart%^&"
	assert.Equal(t, 7, r.GetPoints(receipt))
}

func TestRoundDollarRule(t *testing.T) {
	r := &RoundDollarRule{}
	receipt := getTargetTestReceipt()
	assert.Equal(t, 0, r.GetPoints(receipt))

	receipt.Total = 3500
	assert.Equal(t, 50, r.GetPoints(receipt))
}

func TestRoundQuarterRule(t *testing.T) {
	r := &RoundQuarterRule{}
	receipt := getTargetTestReceipt()
	assert.Equal(t, 0, r.GetPoints(receipt))

	receipt.Total = 2525
	assert.Equal(t, 25, r.GetPoints(receipt))
}

func TestItemPairRule(t *testing.T) {
	r := &ItemPairRule{}
	receipt := getTargetTestReceipt()
	assert.Equal(t, 10, r.GetPoints(receipt))

	receipt.Items = append(receipt.Items, domain.Item{
		ShortDescription: "Knorr Creamy Chicken",
		Price:            126,
	})
	assert.Equal(t, 15, r.GetPoints(receipt))
}

func TestItemDescNameRule(t *testing.T) {
	r := &ItemDescNameRule{}
	receipt := getTargetTestReceipt()
	assert.Equal(t, 6, r.GetPoints(receipt))
}

func TestOddPurchaseDayRule(t *testing.T) {
	r := &OddPurchaseDayRule{}
	receipt := getTargetTestReceipt()
	assert.Equal(t, 6, r.GetPoints(receipt))

	receipt.PurchaseTimestamp = time.Date(2022, 1, 2, 13, 1, 0, 0, time.UTC)
	assert.Equal(t, 0, r.GetPoints(receipt))
}

func TestTwoToFourPurchaseTimeRule(t *testing.T) {
	r := &TwoToFourPurchaseTimeRule{}
	receipt := getTargetTestReceipt()
	assert.Equal(t, 0, r.GetPoints(receipt))

	receipt.PurchaseTimestamp = time.Date(2022, 1, 1, 15, 1, 0, 0, time.UTC)
	assert.Equal(t, 10, r.GetPoints(receipt))
}

func getTargetTestReceipt() *domain.Receipt {
	items := []domain.Item{
		{
			ShortDescription: "Mountain Dew 12PK",
			Price:            649,
		},
		{
			ShortDescription: "Emils Cheese Pizza",
			Price:            1225,
		},
		{
			ShortDescription: "Knorr Creamy Chicken",
			Price:            126,
		},
		{
			ShortDescription: "Doritos Nacho Cheese",
			Price:            335,
		},
		{
			ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
			Price:            1200,
		},
	}

	timestamp, err := time.Parse("2006-01-02 15:04", "2022-01-01 13:01")
	if err != nil {
		panic(err)
	}

	return &domain.Receipt{
		Retailer:          "Target",
		PurchaseTimestamp: timestamp,
		Items:             items,
		Total:             3535,
	}
}
