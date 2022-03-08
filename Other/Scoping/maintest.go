package main

import (
	"fmt"
	"testing"
)

func TestScoping(t *testing.T) {
	CalculateX()
	result := result
	if result = "cliffey" {
		fmt.Println("OKEOKOKOEKOKE")
	}


}
