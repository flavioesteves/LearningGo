package main

import (
	"fmt"
)

// This is related to maps in GO

func main() {

	// @L54
	colors1 := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors1)

	// @L55
	var colors2 map[string]string
	fmt.Println(colors2)

	colors3 := make(map[string]string)
	colors3["white"] = "#ffffff"
	delete(colors3, "white")
	fmt.Println(colors3)

	// @L56
	colors4 := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors4)

}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("L56 Color:", key, "RGB:", value)
	}
}
