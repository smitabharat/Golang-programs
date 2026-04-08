
//i/p- [1,2,3,4]-> [4,3,2,1]
func reversedCopy(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = arr[n-1-i]
	}
	return result
}
