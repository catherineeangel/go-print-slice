package slice

import "fmt"

func Print1DSlice (slice []int) {
	for _, val := range slice{
		fmt.Print(val, " ")
	}
	fmt.Println()
}


func Print2DSlice (slice [][]int) {
	for _, row:= range slice{
		for _, val:=range row{
			fmt.Print(val, "\t")
		}
		fmt.Println()
	}
}