package main

import (
	"fmt"
)

func swap(integer *[]int, index int) {
	if index < 0 || index >= len(*integer)-1 {
		return
	}
	(*integer)[index], (*integer)[index+1] = (*integer)[index+1], (*integer)[index]
}

func BubbleSort() {
	list := make([]int, 0)
	var a int
	fmt.Print("Enter number :")
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		list = append(list, a)
	}
	for pass := 0; pass < len(list)-1; pass++ {
		for i := 0; i < len(list)-1-pass; i++ {
			if list[i] > list[i+1] {
				swap(&list, i)
			}
		}
	}
	for _, v := range list {
		fmt.Printf("%d", v)
	}

}

func main() {

	BubbleSort()

	return

}
