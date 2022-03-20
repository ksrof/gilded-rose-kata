package tests

import (
	"gilded-rose/internal"
	"gilded-rose/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplex(t *testing.T) {
	// Test Backstage passes.
	for day := 0; day < 1; day++ {
		var complexItems = []*models.Item{
			{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 30},
		}

		updatedItem := internal.UpdateComplex(complexItems)

		assert.Equal(t, "Backstage passes to a TAFKAL80ETC concert", updatedItem[0].Name, "name should be equal")
		assert.Equal(t, 14, updatedItem[0].SellIn, "sellIn should decrease by one")
		assert.Equal(t, 31, updatedItem[0].Quality, "quality should increase by one")
	}

	// Test Aged Brie.
	for day := 0; day < 1; day++ {
		var complexItems = []*models.Item{
			{Name: "Aged Brie", SellIn: 10, Quality: 8},
		}

		updatedItem := internal.UpdateComplex(complexItems)

		assert.Equal(t, "Aged Brie", updatedItem[0].Name, "name should be equal")
		assert.Equal(t, 9, updatedItem[0].SellIn, "sellIn should decrease by one")
		assert.Equal(t, 9, updatedItem[0].Quality, "quality should increase by one")
	}

	// Test Sulfuras, Hand of Ragnaros.
	for day := 0; day < 1; day++ {
		var complexItems = []*models.Item{
			{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 80},
		}

		updatedItem := internal.UpdateComplex(complexItems)

		assert.Equal(t, "Sulfuras, Hand of Ragnaros", updatedItem[0].Name, "name should be equal")
		assert.Equal(t, 0, updatedItem[0].SellIn, "sellIn should not decrease")
		assert.Equal(t, 80, updatedItem[0].Quality, "quality should not increase or decrease")
	}
}
