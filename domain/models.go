package domain

import (
	"fmt"
	"time"
)

type Receipt struct {
	Retailer          string
	PurchaseTimestamp time.Time
	Items             []Item
	Total             int
}

func (r *Receipt) ToReceiptDto() *ReceiptDto {
	items := make([]ItemDto, len(r.Items))
	for i, item := range r.Items {
		items[i] = *item.ToItemDto()
	}

	return &ReceiptDto{
		Retailer:     r.Retailer,
		PurchaseDate: r.PurchaseTimestamp.Format("2006-01-02"),
		PurchaseTime: r.PurchaseTimestamp.Format("15:04"),
		Items:        items,
		Total:        FormatMoney(r.Total),
	}
}

type Item struct {
	ShortDescription string
	Price            int
}

func (i *Item) ToItemDto() *ItemDto {
	return &ItemDto{
		ShortDescription: i.ShortDescription,
		Price:            FormatMoney(i.Price),
	}
}

func FormatMoney(money int) string {
	return fmt.Sprintf("%d.%02d", money/100, money%100)
}
