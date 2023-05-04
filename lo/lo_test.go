package lo

import (
	"fmt"
	"github.com/samber/lo"
	"testing"
)

/**
这个库使用泛型实现了大部分的函数式编程函数。不是使用的反射，性能更高，更加安全。
go get github.com/samber/lo@v1
*/

func TestFilter(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := lo.Filter(a, func(x int, index int) bool {
		return x%2 == 0
	})
	fmt.Println(b)
}

func TestMap(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := lo.Map(a, func(item int, index int) string {
		return fmt.Sprintf("%d-%d", index, item)
	})
	fmt.Println(b)
}

func TestFilterMap(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := lo.FilterMap(a, func(item int, index int) (string, bool) {
		if item%2 == 0 {
			return fmt.Sprintf("%d-%d", index, item), true
		} else {
			return "", false
		}
	})
	fmt.Println(b)
}

func TestFlatMap(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := lo.FlatMap(a, func(item int, index int) []string {
		if item%2 == 0 {
			return []string{fmt.Sprintf("%d-%d", index, item), fmt.Sprintf("%d-%d", index, item)}
		} else {
			return []string{}
		}
	})
	fmt.Println(b)
}

func TestUniq(t *testing.T) {
	a := []int{1, 2, 2, 3, 3, 4, 5}
	b := lo.Uniq(a)
	fmt.Println(b)
}

func TestGroupBy(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := lo.GroupBy(a, func(item int) int {
		return item % 2
	})
	fmt.Println(b)
}

func TestReverse(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := lo.Reverse(a)
	fmt.Println(b)
	fmt.Println(a) //原地修改
}

func TestFind(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	e, ok := lo.Find(a, func(item int) bool {
		return item == 3
	})
	fmt.Println(e, ok)
	e, ok = lo.Find(a, func(item int) bool {
		return item == 13
	})
	fmt.Println(e, ok)
}

func TestFindOrElse(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	e := lo.FindOrElse(a, 10, func(item int) bool {
		return item == 3
	})
	fmt.Println(e)
	e = lo.FindOrElse(a, 10, func(item int) bool {
		return item == 13
	})
	fmt.Println(e)
}

func TestEveryBy(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	ok := lo.EveryBy(a, func(item int) bool {
		return item > 0
	})
	fmt.Println(ok)
	ok = lo.EveryBy(a, func(item int) bool {
		return item > 2
	})
	fmt.Println(ok)
}

func TestSomeBy(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	ok := lo.SomeBy(a, func(item int) bool {
		return item > 0
	})
	fmt.Println(ok)
	ok = lo.SomeBy(a, func(item int) bool {
		return item > 2
	})
	fmt.Println(ok)
	ok = lo.SomeBy(a, func(item int) bool {
		return item > 20
	})
	fmt.Println(ok)
}

func TestMin(t *testing.T) {
	min := lo.Min([]int{1, 2, 3})
	fmt.Println(min)

	min = lo.Min([]int{})
	fmt.Println(min)
}

func TestMinBy(t *testing.T) {
	min := lo.MinBy([]string{"s1", "string2", "s3"}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	fmt.Println(min)

	min = lo.MinBy([]string{}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	fmt.Println(min)
}
