package main

import (
	"fmt"
	"sort"
	"sync"
)

func sortArray(a []int, wg *sync.WaitGroup) {
	defer wg.Done()
	n := len(a)
	fmt.Println(a)

	for k := 0; k < n; k++ {
		swapped := false
		for i := 0; i < n-1-k; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	var array [12]int
	var wg sync.WaitGroup

	fmt.Print("Enter the array")
	for i := 0; i < 12; i++ {
		fmt.Scan(&array[i])
	}
	counter := 12 / 4
	k := 4
	index := 0

	for i := 0; i < counter; i++ {
		wg.Add(1)
		go sortArray(array[index:k], &wg)
		index = k
		k += 4
	}
	wg.Wait()

	sort.Ints(array[:])
	fmt.Print(array)

}
