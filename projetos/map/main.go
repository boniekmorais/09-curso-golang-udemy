package main

import "fmt"

func main() {
	fmt.Println("Hello")

	// Map
	// Key => Type string
	// Value => Type string

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	var newColors map[string]string

	colorsFromMake := make(map[string]string)

	colorsFromMake["white"] = "#ffffff"

	colorsIntKey := make(map[int]string)

	colorsIntKey[10] = "#ffffff"
	colorsIntKey[20] = "#34f566"

	delete(colorsIntKey, 10)

	fmt.Println(colors)
	fmt.Println(newColors)
	fmt.Println(colorsFromMake)
	fmt.Println(colorsIntKey)

	printMap(colors)

}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Printf("Color: %s\tHex Code: %s\n", color, hex)
	}
}
