package main

import (
	"fmt"
	"sort"
	"sync"
)

func sortHelp(arrP []int, wg *sync.WaitGroup) {
	sort.Ints(arrP)
	fmt.Println("Sorted Partition ", arrP)
	wg.Done()
}

func main() {
	var size int
	fmt.Printf("Enter array size : ")
	fmt.Scan(&size)

	if size < 4 {
		fmt.Println("Minimum size must be 4")
		return
	}

	var arr []int
	fmt.Println("Enter Elements")
	for i := 0; i < size; i++ {
		var val int
		fmt.Scan(&val)
		arr = append(arr, val)
	}

	partitionSize := size / 4

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		if i != 3 {
			go sortHelp(arr[i*partitionSize:(i+1)*partitionSize], &wg)
		} else {
			go sortHelp(arr[i*partitionSize:], &wg)
		}
	}

	wg.Wait()

	p1 := arr[0:partitionSize]
	p2 := arr[partitionSize*1 : partitionSize*2]
	p3 := arr[partitionSize*2 : partitionSize*3]
	p4 := arr[partitionSize*3:]

	i, j, k, l, min, index := 0, 0, 0, 0, 0, 0

	var sortedArr []int

	for m := 0; m < len(arr); m++ {
		min = 2147483647
		index = 0

		if i < len(p1) {
			min = p1[i]
			index = 1
		}
		if j < len(p2) {
			if min > p2[j] {
				min = p2[j]
				index = 2
			}
		}
		if k < len(p2) {
			if min > p3[k] {
				min = p3[k]
				index = 3
			}
		}
		if l < len(p4) {
			if min > p4[l] {
				min = p4[l]
				index = 4
			}
		}

		sortedArr = append(sortedArr, min)

		if index == 1 {
			i++
		} else if index == 2 {
			j++
		} else if index == 3 {
			k++
		} else if index == 4 {
			l++
		}

	}

	fmt.Println("Sorted Array ", sortedArr)

}
