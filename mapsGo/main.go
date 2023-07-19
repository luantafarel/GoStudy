package main

import "fmt"

func main() {
	// colors := make(map[int]string)

	// var colors map[string]string

	colors := map[string]string{
		"red":  "FF0000",
		"blue": "0000FF",
	}

	// colors[10] = "#ffffff"
	// colors[1] = "#ffffff"
	// colors[2] = "#ffffff"
	// delete(colors, 10)

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key)
		fmt.Println(value)
	}
}
