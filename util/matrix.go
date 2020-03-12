package util

import "fmt"

func PrintIntMatrix(nodes [][]int) {
	for _, row := range nodes {
		fmt.Println(row)
	}
	fmt.Println("----------")
}
