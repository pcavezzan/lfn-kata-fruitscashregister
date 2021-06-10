package app

import "fmt"

type App struct {

}

func New() *App {
	return &App{}
}


func(app *App) Run() {
	cart := newShoppingCart()
	for {
		var name string
		fmt.Scanf("%s\n", &name)
		cart.add(name)
		fmt.Println(cart.printStatus())
	}
}