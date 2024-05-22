package array

func reverseArray(arr []int) []int {

	last := len(arr)
	mid := last / 2

	for i := 0; i < mid; i++ {

		arr[i], arr[mid-i] = arr[mid-i], arr[i]
	}
	return arr
}
