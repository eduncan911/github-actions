package main

import (
	"fmt"
)

// Thing is a thing
type Thing struct {
	Thang string
}

// ThingThang takes in an input, and combines it with Thang
func (t Thing) ThingThang(thing string) string {
	if thing == "this" {
		return t.Thang + thing
	} else if thing == "that" {
		return thing + t.Thang
	}
	return thing
}

func main() {
	t := Thing{Thang: "This"}

	fmt.Println(t.ThingThang("That"))
}
