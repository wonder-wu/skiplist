package main

import (
	"fmt"
	"github.com/wonder-wu/skiplist"
	"strconv"
)

type Number int

func (n Number) CompareTo(i interface{}) int {
	number, ok := i.(Number)
	if !ok {
		panic("need int")
	}
	if number > n {
		return -1
	} else if number == n {
		return 0
	} else {
		return 1
	}
}
func (n Number) String() string {
	return strconv.Itoa(int(n))
}

func main() {
	skp := skiplist.New(8)

	for i := 0; i < 50; i++ {
		skp.Insert(Number(i), i)
	}

	fmt.Println("structure:")
	skp.PrintStructure()

	fmt.Println("search 12:")
	val, ok := skp.Search(Number(12))
	fmt.Println(val, ok)

	fmt.Println("search 10:")
	val, ok = skp.Search(Number(10))
	fmt.Println(val, ok)

	fmt.Println("delete 10:")
	skp.Delete(Number(10))
	fmt.Println("search 10:")
	val, ok = skp.Search(Number(10))
	fmt.Println(val, ok)
	fmt.Println("structure:")
	skp.PrintStructure()
}
