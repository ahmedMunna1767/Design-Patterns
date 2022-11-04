package main

import "fmt"

func main() {

	pizza := &VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggieMania with %v is %d\n", pizzaWithCheeseAndTomato.getComponents(), pizzaWithCheeseAndTomato.getPrice())
}
