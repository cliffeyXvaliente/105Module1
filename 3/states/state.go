package state

import "fmt"

// A string Array for the game
var West = [4]string{"Homosapiens", "Fox", "Chicken", "Korn"}

// Make variables taking the values of the array
var ggetHuman string = West[0]
var ggetFox string = West[1]

// Beginning starting

func ViewState() {

	fmt.Printf("======================================== STARTING HERE ======================================== \n")

	// We print the arrays West

	fmt.Printf("%v -----\\=======//________________________________________________----------  \n", West)

}

// RiverCrossing Function

func RiverCrossingX() {

	fmt.Printf("======================================== ON THE RIVER ========================================\n")

	// We make variables to get each value of an array

	// We make a new variable that merge or plus the values of the variables we made from the Slice or Array global variables*
	var LanstateItems string = ggetHuman + " " + ggetFox

	fmt.Printf("----____________________________\\===%v====//_________________________-----------\n", LanstateItems)

}

func FinalEnd() {

	fmt.Printf("======================================== ON THE END ========================================\n")

	// We print both on the land
	fmt.Printf("----_____________________________________________________\\===%v %v====//-----------\n", ggetHuman, ggetFox)

}
