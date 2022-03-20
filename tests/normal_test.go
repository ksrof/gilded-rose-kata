package tests

import (
	"gilded-rose/internal"
	"gilded-rose/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormal(t *testing.T) {
	// Test Normal.
	for day := 0; day < 1; day++ {
		var normalItems = []*models.Item{
			{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 20},
		}

		updatedItem := internal.UpdateNormal(normalItems)

		assert.Equal(t, "+5 Dexterity Vest", updatedItem[0].Name, "name should be equal")
		assert.Equal(t, 9, updatedItem[0].SellIn, "sellIn should decrease by one")
		assert.Equal(t, 19, updatedItem[0].Quality, "quality should decrease by one")
	}
}
