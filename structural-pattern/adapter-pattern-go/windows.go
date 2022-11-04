package main

import "fmt"

//  concrete windows doesn't implements computer
type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
