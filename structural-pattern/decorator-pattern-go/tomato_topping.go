package main

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

func (c *TomatoTopping) getComponents() []string {
	pizzaComponents := c.pizza.getComponents()
	return append(pizzaComponents, "tomato slices")
}
