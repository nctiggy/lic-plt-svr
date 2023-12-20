// Package api manages the api interaction for manipulating and deleting
// New license plates
package api

import "fmt"

type licensePlate struct {
	country string
	state   string
	img     licPltImg
}

type licPltImg struct {
	plateName string
	imgURL    string
}

func main() {
	fmt.Println("vim-go")
}
