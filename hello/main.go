package main

import (
	"fmt"
	"log"
	"math"

	"example.com/greetings"
)

func needInt(x int) int {
	return x * 10
}

func needFloat(x float64) float64 {
	return x * 10
}

const (
	Big   = 1 << 16
	Small = Big >> 15
)

func main() {
	log.SetPrefix("greetings log: ")
	log.SetFlags(0)
	messages, err := greetings.Hellos([]string{"Tud", "Lian"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

	a := 2
	b := 3
	result := math.Sqrt(float64(a*2 + b*2))
	fmt.Printf("result is %v", result)

	fmt.Printf("needInt(Small): %v\n", needInt(Small))
	fmt.Printf("needInt(Big): %v\n", needInt(Big))
	fmt.Printf("needFloat(Small): %v\n", needFloat(Small))
	fmt.Printf("needFloat(Big): %v\n", needFloat(Big))

}
