package main

import (
	"fmt"
	"reflect"
)

type SomeStruct struct {
	yo int
}

func main() {
	first := 2
	second := "second"
	third := true
	fourth := 4.5
	fifth := SomeStruct{
		yo: 5,
	}

	fmt.Printf(
		"first var is a %s\n"+
			"second var is a %s\n"+
			"third var is a %s\n"+
			"fourth var is a %s\n"+
			"fifth var is a %s\n",
		getType(first).String(),
		getType(second).String(),
		getType(third).String(),
		getType(fourth).String(),
		getType(fifth).String(),
	)
}

func getType(obj interface{}) reflect.Type {
	return reflect.TypeOf(obj)
}
