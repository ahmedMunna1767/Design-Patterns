package main

import "fmt"

type musket struct {
	Gun
}

// not accessible from the gun interface
func (m *musket) Shot() {
	fmt.Println("thus")
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
