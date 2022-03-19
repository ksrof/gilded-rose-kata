package main

import (
	"fmt"
	"gilded-rose/internal"
	"gilded-rose/models"
)

// update calls different methods that modify the values of the items based on their names.
// It also prints the updated items separated by categories.
func update(items []*models.Item) {
	fmt.Println("\n------ Complex ------")
	internal.Complex(items)

	fmt.Println("\n------ Conjured ------")
	internal.Conjured(items)

	fmt.Println("\n------ Normal ------")
	internal.Normal(items)
}

func main() {
	var items = []*models.Item{
		{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 20},
		{Name: "Aged Brie", SellIn: 2, Quality: 0},
		{Name: "Elixir of the Mongoose", SellIn: 5, Quality: 7},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 80},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
		{Name: "Conjured Mana Cake", SellIn: 3, Quality: 6},
	}

	for day := 0; day < 5; day++ {
		fmt.Printf("\n####### Day %d #######", day)
		update(items)
	}
}
