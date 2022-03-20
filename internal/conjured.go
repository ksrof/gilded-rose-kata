package internal

import (
	"fmt"
	"gilded-rose/models"
)

// UpdateConjured decreases the quality of an item based on a series of scenarios.
// It decreases the quality by two each day, and by four when time expired.
func UpdateConjured(items []*models.Item) []*models.Item {
	for i := 0; i < len(items); i++ {
		items[i].SellIn = items[i].SellIn - 1

		if items[i].SellIn >= 0 && items[i].Quality > 0 {
			items[i].Quality = items[i].Quality - 2
		}

		if items[i].SellIn <= 0 && items[i].Quality > 0 {
			items[i].Quality = items[i].Quality - 4
		}

		fmt.Printf("Name: %s, SellIn: %d, Quality: %d\n", items[i].Name, items[i].SellIn, items[i].Quality)
	}

	return items
}
