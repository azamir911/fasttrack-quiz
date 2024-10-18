package main

import "fmt"

func GetTwoOnSum(nums []int, target int) []int {
	indices := make(map[int]int)

	for i, num := range nums {
		diff := target - num

		if iDiff, ok := indices[diff]; ok {
			return []int{iDiff, i}
		}

		indices[num] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	result := GetTwoOnSum(nums, 17)
	fmt.Println("The result is: ", result)
}
