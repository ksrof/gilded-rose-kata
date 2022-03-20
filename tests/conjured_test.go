package tests

import (
	"gilded-rose/internal"
	"gilded-rose/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConjured(t *testing.T) {
	// Test Conjured.
	for day := 0; day < 1; day++ {
		var conjuredItems = []*models.Item{
			{Name: "Conjured Mana Cake", SellIn: 3, Quality: 6},
		}

		updatedItem := internal.UpdateConjured(conjuredItems)

		assert.Equal(t, "Conjured Mana Cake", updatedItem[0].Name, "name should be equal")
		assert.Equal(t, 2, updatedItem[0].SellIn, "sellIn should decrease by one")
		assert.Equal(t, 4, updatedItem[0].Quality, "quality should decrease by two")
	}
}
