package main


func rob(nums []int) int {
	prev1 := 0
	prev2 := 0
	curr := 0
	for _, num := range nums {
		curr = max(prev2+num, prev1)
		prev2 = prev1
		prev1 = curr
	}
	return curr
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

