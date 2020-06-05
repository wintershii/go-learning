package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1,2,3},{4,5,6},{7,8,9}}))
}


func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := make([]int,len(matrix) * len(matrix[0]))
	l := (Min(len(matrix),len(matrix[0])) + 1) / 2
	k := 0
	for i := 0; i < l  && k < len(res); i++ {
		for j := i; j < len(matrix[0]) - i && k < len(res); j++ {
			res[k] = matrix[i][j]
			k++
		}
		for j := i+1; j < len(matrix) - i  && k < len(res); j++ {
			res[k] = matrix[j][len(matrix[0])-i-1]
			k++
		}
		for j := len(matrix[0])-i-2; j >= i  && k < len(res); j-- {
			res[k] = matrix[len(matrix)-i-1][j]
			k++
		}
		for j := len(matrix) - i-2; j > i  && k < len(res); j-- {
			res[k] = matrix[j][i]
			k++
		}
	}
	return res
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}