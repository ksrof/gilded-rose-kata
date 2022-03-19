package internal

import (
	"fmt"
	"gilded-rose/models"
)

// Complex increases the quality of an item based on a series of scenarios.
// The incrementation/decrementation of the quality is based on the name of the item.
func Complex(items []*models.Item) []*models.Item {
	for i := 0; i < len(items); i++ {
		// Backstage passes quality increments every day by one.
		// When the sell time is less than or equal to ten quality increments by two.
		// When the sell time is less than or equal to five quality increments by three.
		// When the sell time is less than or equal to zero quality downs to zero.
		if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
			items[i].SellIn = items[i].SellIn - 1

			if items[i].SellIn > 10 && items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
			}

			if items[i].SellIn <= 10 && items[i].SellIn > 5 && items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 2
			}

			if items[i].SellIn <= 5 && items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 3
			}

			if items[i].SellIn <= 0 {
				items[i].Quality = 0
			}

			fmt.Printf("Name: %s, SellIn: %d, Quality: %d\n", items[i].Name, items[i].SellIn, items[i].Quality)
		}

		// Aged Brie quality increments every day by one.
		// When the sell time is less than or equal to zero quality increments by two.
		if items[i].Name == "Aged Brie" {
			items[i].SellIn = items[i].SellIn - 1

			if items[i].SellIn > 0 && items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
			}

			if items[i].SellIn <= 0 && items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 2
			}

			fmt.Printf("Name: %s, SellIn: %d, Quality: %d\n", items[i].Name, items[i].SellIn, items[i].Quality)
		}

		// Sulfuras, Hand of Ragnaros its a legendary article
		// Its sell time and its quality aren't neither modified or degraded.
		// The easiest one by far :D
		if items[i].Name == "Sulfuras, Hand of Ragnaros" {
			fmt.Printf("Name: %s, SellIn: %d, Quality: %d\n", items[i].Name, items[i].SellIn, items[i].Quality)
		}
	}

	return items
}
