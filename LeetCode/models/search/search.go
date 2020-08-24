package search

func BinarySreach(arr []int, key int) int {
	res := -1
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		temp := arr[mid]
		if temp == key {
			res = mid
			break
		} else if temp > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return res
}
