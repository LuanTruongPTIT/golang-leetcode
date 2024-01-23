package main

import (
	"fmt"
	"strings"
)

var n int = 10

func swap(x, y string) (string, string) {
	var i string
	i = x
	x = y
	y = i
	return x, y
}

type Point struct {
	X, Y float64
}
type ColoredPoint struct {
	Point
}

func main() {
	var k float32
	k = 62.4
	var j int
	j = 64
	var l int = 66
	i := 11
	fmt.Println(i, j, l)
	fmt.Printf("%T, %v", i, i)
	fmt.Println(k, n)
	a, b := swap("luantruong", "truongluan")
	fmt.Println(a, b)
	greeting := []string{"Hello", "World"}
	fmt.Println(strings.Join(greeting, ","))
	// var c = [...]int{2: 3, 1: 2}
	var d = [...]int{1, 2, 4: 5, 6}

	for index, value := range d {

		fmt.Println(index, value)
	}
	fmt.Println("--------------------")
	var m = []int{1, 3, 4, 5}
	// var n = &m
	// l := make()
	stack, res := []int{m[0]}, 0
	fmt.Println(stack, res)
	// for index, value := range n {
	// 	fmt.Println(index, value)
	// }

}
