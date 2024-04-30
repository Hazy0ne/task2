package main

import "fmt"

func main() {
	arr := [4]int{12, 30, 48, 24}
	slise := []int{}
	sliseCop := []int{}
	m := arr[0]
	for i := 0; i < len(arr); i++ {
		if m > arr[i] {
			m = arr[i]
		}
	}

	for i := 1; i <= m; i++ {
		slise = append(slise, i)
	}

	for i := 0; i < len(arr); i++ {
		for b := 0; b < m; b++ {
			if arr[i]%(b+1) != 0 {
				slise[b] = 0
			}
		}
	}

	for i := 0; i < len(slise); i++ {
		if slise[i] > 1 {
			sliseCop = append(sliseCop, slise[i])
		}
	}

	fmt.Println(sliseCop)
}
