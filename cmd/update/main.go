package main

import (
	"fmt"
	"gilded-rose/internal"
	"gilded-rose/models"
)

func main() {
	var complexItems = []*models.Item{
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
		{Name: "Aged Brie", SellIn: 2, Quality: 0},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 80},
	}

	var conjuredItems = []*models.Item{
		{Name: "Conjured Mana Cake", SellIn: 3, Quality: 6},
	}

	var normalItems = []*models.Item{
		{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 20},
		{Name: "Elixir of the Mongoose", SellIn: 5, Quality: 7},
	}

	// Update the items each day.
	for day := 0; day < 5; day++ {
		// Print the day.
		fmt.Printf("\n-------- Day %d --------\n", day)
		internal.UpdateComplex(complexItems)
		internal.UpdateConjured(conjuredItems)
		internal.UpdateNormal(normalItems)
	}
}
