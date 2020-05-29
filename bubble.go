package main

import "fmt"

func BubbleSort(sli []int){
	for i:=0; i<len(sli); i++{
		for j:=i; j<len(sli); j++{
			if(sli[i]>sli[j]){
				Swap(i, j, sli)
			}
		}
	}

}

func Swap(i int, j int, sli []int){
	temp:= sli[i]
	sli[i] = sli[j]
	sli[j] = temp

}

func main() {
	fmt.Println("Please type in 10 integers")
	var num int
	arr := make([]int, 0, 10)

	for i:=0; i<10; i++{
		fmt.Scan(&num)
		arr = append(arr, num)
	}
	BubbleSort(arr)
	fmt.Println(arr)

}