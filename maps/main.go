package main

import "fmt"

type mymap map[string]string

func main() {
	colors := mymap{
		"red":   "#ff0000",
		"black": "#000000",
		"green": "#00ff00",
	}

	colors.printColors()
}

func (c mymap) printColors() {
	for color, hex := range c {
		fmt.Println(color + " " + hex)
	}
}
