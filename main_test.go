package main

import (
	"fmt"
	"testing"
)

func TestThis(t *testing.T) {
	// arrange
	something := "something"
	thing := Thing{}

	// act
	v := thing.ThingThang(something)

	// assert
	if v != something {
		fmt.Printf("TestThingThang failed. Wanted %s got %s", something, v)
		t.Fail()
	}
}

// func TestThisThat(t *testing.T) {
// 	// arrange
// 	this := "this"
// 	that := "that"
// 	thing := Thing{Thang: that}

// 	// act
// 	v := thing.ThingThang(this)

// 	// assert
// 	expected := that + this
// 	if v != expected {
// 		fmt.Printf("TestThingThang failed. Wanted %s got %s", expected, v)
// 		t.Fail()
// 	}
// }

func TestThatThis(t *testing.T) {
	// arrange
	this := "this"
	that := "that"
	thing := Thing{Thang: this}

	// act
	v := thing.ThingThang(that)

	// assert
	expected := that + this
	if v != expected {
		fmt.Printf("TestThingThang failed. Wanted %s got %s", expected, v)
		t.Fail()
	}
}
