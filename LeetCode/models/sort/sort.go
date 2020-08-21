package sort

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

// O(n^2)
func BubbleSort(nums *[]int) {
	size := len(*nums) - 1
	for i := 0; i <= size; i++ {
		for j := size; j > i; j-- {
			if (*nums)[j] < (*nums)[j-1] {
				swap(nums, j-1, j)
			}
		}
	}
}

func BubbleSortV2(nums *[]int) {
	size := len(*nums) - 1
	flag := true
	for i := 0; i <= size && flag; i++ {
		flag = false
		for j := size; j > i; j-- {
			if (*nums)[j] < (*nums)[j-1] {
				swap(nums, j-1, j)
				flag = true
			}
		}
	}
}

func SelectSort(nums *[]int) {
	size := len(*nums) - 1
	for i := 0; i <= size; i++ {
		min := (*nums)[size]
		minIndex := size
		for j := size; j > i; j-- {
			if (*nums)[j] < min {
				min = (*nums)[j]
				minIndex = j
			}
		}
		swap(nums, i, minIndex)
	}
}

func SelectSortV2(nums *[]int) {
	size := len(*nums) - 1
	for i := 0; i <= size; i++ {
		min := i
		for j := size; j > i; j-- {
			if (*nums)[j] < (*nums)[min] {
				min = j
			}
		}
		swap(nums, i, min)
	}
}

func InsertSort(nums *[]int) {
	size := len(*nums) - 1
	for i := 1; i <= size; i++ {
		j := i
		for j > 0 {
			if (*nums)[j-1] > (*nums)[j] {
				swap(nums, j-1, j)
			}
			j--
		}
	}
}

func InsertSortV2(nums *[]int) {
	size := len(*nums) - 1
	for i := 1; i <= size; i++ {
		j := i
		flag := true
		for j > 0 && flag {
			flag = false
			if (*nums)[j-1] > (*nums)[j] {
				swap(nums, j-1, j)
				flag = true
			}
			j--
		}
	}
}

func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func Merge(left, right []int) []int {
	k, i, j := len(left)+len(right), 0, 0
	slice := make([]int, k)
	for m := 0; m < k; m++ {
		if i >= len(left) && j < len(right) {
			//left用完
			slice[m] = right[j]
			j++
		} else if i < len(left) && j >= len(right) {
			//right
			slice[m] = left[i]
			i++
		} else {
			if left[i] < right[j] {
				slice[m] = left[i]
				i++
			} else {
				slice[m] = right[j]
				j++
			}
		}

	}
	return slice
}

func QuickSort(nums *[]int, low, high int) {
	if low < high {
		pivot := Partition(nums, low, high)
		QuickSort(nums, 0, pivot-1)
		QuickSort(nums, pivot+1, high)
	}
}

func Partition(nums *[]int, low, high int) int {
	pivot := high
	temp := (*nums)[pivot]
	// fmt.Printf("pivot:%v,temp:%v\n", pivot, temp)
	for low < high {
		for (*nums)[low] < temp {
			low++
		}
		swap(nums, low, high)

		for (*nums)[high] > temp {
			high--
		}
		swap(nums, low, high)
	}
	return low
}

// undo//:heap sort
