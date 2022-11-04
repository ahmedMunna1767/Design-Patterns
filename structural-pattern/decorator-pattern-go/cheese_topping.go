package main

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

func (c *CheeseTopping) getComponents() []string {
	pizzaComponents := c.pizza.getComponents()
	return append(pizzaComponents, "cheese slices")
}
