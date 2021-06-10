package app

import "fmt"

type catalog struct {
	products []product
}

type shoppingCart struct {
	c            catalog
	productNames []product
	total 		int
}

func newShoppingCart() *shoppingCart {
	return &shoppingCart{
		c: catalog{products: []product{
			{
				Name:     "Pommes",
				Price:    100,
				Discount: 0,
			},
			{
				Name:     "Bananes",
				Price:    150,
				Discount: 0,
				NumberForFree: 2,
			},
			{
				Name:     "Cerises",
				Price:    75,
				Discount: 20,
				NumberForDiscount: 2,
			},
			{
				Name:     "Apples",
				Price:    100,
				Discount: 0,
			},
		}},
	}
}

func (sc *shoppingCart) getNumberOfSimilarProduct(product product) int {
	var nb int
	for _, p := range sc.productNames {
		if p.Name == product.Name {
			nb++
		}
	}
	return nb
}

func (sc *shoppingCart) add(p string) {
	productFound, found := find(sc.c.products, p)
	if found {
		sc.productNames = append(sc.productNames, *productFound)
		for _, products := range sc.productNames {
			sc.total += products.getPrice(sc.getNumberOfSimilarProduct(products))
		}
	}
}


func (sc *shoppingCart) printStatus() string {
	return fmt.Sprintf("%s > %d\n", sc.productNames[len(sc.productNames) -1].Name, sc.total)
}

type product struct {
	Name  string
	Price int
	Discount int
	NumberForDiscount int
	NumberForFree int
}

func (p product) getPrice(n int) int {
	if p.NumberForFree != 0 && n >= p.NumberForFree {
		return 0
	}
	if n >= p.NumberForDiscount {
		return p.Price - p.Discount
	}
	return p.Price
}

// find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func find(slice []product, val string) (*product, bool) {
	for _, item := range slice {
		if item.Name == val {
			return &item, true
		}
	}
	return nil, false
}