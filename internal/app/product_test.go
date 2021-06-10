package app

import "testing"
import "github.com/stretchr/testify/assert"

func Test_shoppingCart_add_shouldAddAProduct(t *testing.T) {
	s := newShoppingCart()

	s.add("Pommes")

	assert.Len(t, s.productNames, 1)
}

func Test_shoppingCart_total_shouldReturnPriceOfShoppingCard(t *testing.T) {
	s := newShoppingCart()
	s.add("Pommes")

	total := s.total
	fruits := s.productNames

	assert.EqualValues(t, 100, total)
	assert.EqualValues(t, []product{{
		Name:     "Pommes",
		Price:    100,
		Discount: 0,
	}}, fruits)
}

func Test_shoppingCart_total_shouldReturnSupportInternationalization(t *testing.T) {
	s := newShoppingCart()
	s.add("Apples")

	total := s.total
	fruits := s.productNames

	assert.EqualValues(t, 100, total)
	assert.EqualValues(t, []product{{
		Name:     "Apples",
		Price:    100,
		Discount: 0,
	}}, fruits)
}

func Test_shoppingCart_total_shouldReturnPriceWithDiscount(t *testing.T) {
	s := newShoppingCart()
	s.add("Cerises")
	s.add("Cerises")

	total := s.total
	fruits := s.productNames

	assert.EqualValues(t, 185, total)
	assert.EqualValues(t, []product{
		{
			Name:     "Cerises",
			Price:    75,
			Discount: 20,
			NumberForDiscount: 2,
		},
		{
			Name:     "Cerises",
			Price:    75,
			Discount: 20,
			NumberForDiscount: 2,
		},
	}, fruits)
}

func Test_shoppingCart_total_shouldReturnPriceWithFree(t *testing.T) {
	s := newShoppingCart()
	s.add("Bananes")
	s.add("Bananes")

	total := s.total
	fruits := s.productNames

	assert.EqualValues(t, 150, total)
	assert.EqualValues(t, []product{
		{
			Name:     "Bananes",
			Price:    150,
			Discount: 0,
			NumberForFree: 2,
		},
		{
			Name:     "Bananes",
			Price:    150,
			Discount: 0,
			NumberForFree: 2,
		},
	}, fruits)
}