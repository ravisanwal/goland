package main

import (
	"github.com/ravisanwal/goland/structs"
	"reflect"
	"fmt"
)


func mains() {
	one := structs.One{Value: "Hello"}
	two := structs.Two{AnotherValue: "World!"}
	fmt.Println(one)
	fmt.Println(two)

	fmt.Println(reflect.TypeOf(one))
	fmt.Println(reflect.TypeOf(two))
}