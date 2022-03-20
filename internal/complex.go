package internal

import (
	"fmt"
	"gilded-rose/models"
)

// UpdateComplex increases the quality of an item based on a series of scenarios.
// The incrementation/decrementation of the quality is based on the name of the item.
func UpdateComplex(items []*models.Item) []*models.Item {
	for i := 0; i < len(items); i++ {
		switch items[i].Name {
		case "Backstage passes to a TAFKAL80ETC concert":
			items[i].SellIn = items[i].SellIn - 1

			// Backstage passes quality increments every day by one.
			if items[i].SellIn > 10 && items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
			}

			// When the sell time is less than or equal to ten quality increments by two.
			if items[i].SellIn <= 10 && items[i].SellIn > 5 && items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 2
			}

			// When the sell time is less than or equal to five quality increments by three.
			if items[i].SellIn <= 5 && items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 3
			}

			// When the sell time is less than or equal to zero quality downs to zero.
			if items[i].SellIn <= 0 {
				items[i].Quality = 0
			}

			fmt.Printf("Name: %s, SellIn: %d, Quality: %d\n", items[i].Name, items[i].SellIn, items[i].Quality)
		case "Aged Brie":
			items[i].SellIn = items[i].SellIn - 1

			// Aged Brie quality increments every day by one.
			if items[i].SellIn > 0 && items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
			}

			// When the sell time is less than or equal to zero quality increments by two.
			if items[i].SellIn <= 0 && items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 2
			}

			fmt.Printf("Name: %s, SellIn: %d, Quality: %d\n", items[i].Name, items[i].SellIn, items[i].Quality)
		case "Sulfuras, Hand of Ragnaros":
			// Its sell time and its quality aren't neither modified or degraded.
			// The easiest one by far :D
			fmt.Printf("Name: %s, SellIn: %d, Quality: %d\n", items[i].Name, items[i].SellIn, items[i].Quality)
		}
	}

	return items
}
