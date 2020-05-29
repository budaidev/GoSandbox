package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func sortSlice(slc []int, tname string) {
	fmt.Println("thread-", tname, " will sort", slc)
	sort.Ints(slc)
	fmt.Println("subresult: of thread-", tname, "->", slc)
	defer wg.Done()
}

func convertLineToArray(line string) []int {
	s := strings.Split(line, " ")
	var numbers []int
	for _, v := range s {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func findMin(arr []int, indexes []int) (int, int) {
	minIndex := 0
	minValue := arr[indexes[minIndex]]
	for index, i := range indexes {
		if arr[i] < minValue {
			minIndex = index
			minValue = arr[i]
		}

	}
	return minValue, minIndex
}

func removeIndex(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	fmt.Println("Please enter the number you want to sort, separated by space")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbers := convertLineToArray(scanner.Text())

	numberOfRoutines := 4

	sliceSize := len(numbers) / numberOfRoutines

	sliceStarts := make([]int, numberOfRoutines+1)

	for i := 0; i < numberOfRoutines; i++ {
		sliceStarts[i] = i * sliceSize
	}
	sliceStarts[numberOfRoutines] = len(numbers)

	for i := 0; i < numberOfRoutines; i++ {
		wg.Add(1)
		go sortSlice(numbers[sliceStarts[i]:sliceStarts[i+1]], strconv.Itoa(i+1))
	}
	wg.Wait()

	//merge
	result := make([]int, 0)
	sliceIndexes := make([]int, numberOfRoutines)
	copy(sliceIndexes, sliceStarts[:len(sliceStarts)-1])

	for {
		if len(result) >= len(numbers) {
			break
		}
		minValue, minIndex := findMin(numbers, sliceIndexes)

		result = append(result, minValue)
		if sliceIndexes[minIndex]+1 < sliceStarts[minIndex+1] {
			sliceIndexes[minIndex]++
		} else {
			sliceIndexes = removeIndex(sliceIndexes, minIndex)
		}

	}
	fmt.Println("The result is", result)

}
