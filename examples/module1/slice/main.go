package main

import (
	"fmt"
	"reflect"
)

type MyTpye struct {
	Name string `json:"name"`
}

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:3]
	fmt.Printf("mySlice %+v\n", mySlice)
	fullSlice := myArray[:]
	remove3rdItem := deleteItem(fullSlice, 2)
	fmt.Printf("remove3rdItem %+v\n", remove3rdItem)

	mt := MyTpye{Name: "gwgrisk"}
	myType := reflect.TypeOf(mt)
	name := myType.Field(0)
	tag := name.Tag.Get("json")

	fmt.Printf("name %s\n", tag)
}

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
