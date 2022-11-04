package main

type VeggieMania struct {
}

// concrete component
func (p *VeggieMania) getPrice() int {
	return 15
}

func (p *VeggieMania) getComponents() []string {
	return []string{"vegetables"}
}
