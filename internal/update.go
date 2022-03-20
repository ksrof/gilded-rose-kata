package internal

import "gilded-rose/models"

// UpdateComplex modifies the complex items.
// It also prints the updated items.
func UpdateComplex(items []*models.Item) {
	// Update the complex items.
	Complex(items)
}

// UpdateConjured modifies the conjured items.
// It also prints the updated items.
func UpdateConjured(items []*models.Item) {
	// Update the conjured items.
	Conjured(items)
}

// UpdateNormal modifies the normal items.
// It also prints the updated items.
func UpdateNormal(items []*models.Item) {
	// Update the normal items.
	Normal(items)
}
