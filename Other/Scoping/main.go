package main

import "fmt"

func main() {

	Scoping()
	makeLoop()

}

// New func

func Scoping() {

	Sniper := "Snaipah"
	fmt.Println(Sniper)
}

func makeLoop() {

	i := 20
	y := 21

	if i <= y {

		fmt.Println("gg ezzz")
		fmt.Println(i)
	}
}

func CalculateX() {

	var result string = "cliffey"
	println(result)
	return

}
