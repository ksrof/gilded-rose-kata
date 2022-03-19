package internal

import (
	"fmt"
	"gilded-rose/models"
	"strings"
)

// Normal decreases the quality of an item based on a series of scenarios.
// It decreases the quality by one each day, and by two when time expired.
func Normal(items []*models.Item) []*models.Item {
	for i := 0; i < len(items); i++ {
		// Check that complex items do not follow the rules of this method.
		if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" && items[i].Name != "Aged Brie" && items[i].Name != "Sulfuras, Hand of Ragnaros" {
			// Check that conjured items do not follow the rules of this method.
			if !strings.Contains(items[i].Name, "Conjured") {
				items[i].SellIn = items[i].SellIn - 1

				if items[i].SellIn >= 0 && items[i].Quality > 0 {
					items[i].Quality = items[i].Quality - 1
				}

				if items[i].SellIn <= 0 && items[i].Quality > 0 {
					items[i].Quality = items[i].Quality - 2
				}

				fmt.Printf("Name: %s, SellIn: %d, Quality: %d\n", items[i].Name, items[i].SellIn, items[i].Quality)
			}
		}
	}

	return items
}
