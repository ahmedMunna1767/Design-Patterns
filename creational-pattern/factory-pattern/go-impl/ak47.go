package main

import "fmt"

type Ak47 struct {
	Gun
}

// not accessible from the gun interface
func (ak *Ak47) BrushFire() {
	fmt.Println("Grr")
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
