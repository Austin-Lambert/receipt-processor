package domain

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ReceiptDto struct {
	Retailer     string    `json:"retailer"`
	PurchaseDate string    `json:"purchaseDate"`
	PurchaseTime string    `json:"purchaseTime"`
	Items        []ItemDto `json:"items"`
	Total        string    `json:"total"`
}

func (r *ReceiptDto) ToReceipt() (*Receipt, error) {
	items := make([]Item, len(r.Items))
	for i, item := range r.Items {
		model, err := item.ToItem()
		if err != nil {
			panic(err)
		}
		items[i] = *model
	}
	money, err := ParseMoney(r.Total)
	if err != nil {
		return nil, err
	}
	timestamp, err := MergeDateTime(r.PurchaseDate, r.PurchaseTime)
	if err != nil {
		return nil, err
	}
	return &Receipt{
		Retailer:          r.Retailer,
		PurchaseTimestamp: timestamp,
		Items:             items,
		Total:             money,
	}, nil
}

type ItemDto struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

func (i *ItemDto) ToItem() (*Item, error) {
	money, err := ParseMoney(i.Price)
	if err != nil {
		return nil, err
	}
	return &Item{
		ShortDescription: i.ShortDescription,
		Price:            money,
	}, nil
}

func MergeDateTime(dateStr, timeStr string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date: %w", err)
	}

	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time: %w", err)
	}

	merged := time.Date(date.Year(), date.Month(), date.Day(), t.Hour(), t.Minute(), 0, 0, time.UTC)
	return merged, nil
}

func ParseMoney(money string) (int, error) {
	str := strings.ReplaceAll(money, ".", "")
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return num, nil
}
