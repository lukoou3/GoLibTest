package functions

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := Filter(a, func(i int) bool {
		return i%2 == 0
	})
	c := FilterNot(a, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func TestMap(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := Map(a, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(a)
	fmt.Println(b)
}

func TestForEach(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	ForEach(a, func(i int) {
		fmt.Println(i)
	})
}
