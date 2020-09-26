package sort

import "fmt"

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

		for low < high && (*nums)[low] <= temp {
			low++
		}
		swap(nums, low, high)

		for low < high && (*nums)[high] >= temp {
			high--
		}
		swap(nums, low, high)
	}
	return low
}

func PartitionHasErr(nums *[]int, low, high int) int {
	pivot := high
	temp := (*nums)[pivot]
	fmt.Printf("pivot:%v,temp:%v\n", pivot, temp)
	for low < high {
		// 沒有比較條件就會往上飄到爆range
		for (*nums)[low] <= temp {
			low++
		}
		swap(nums, low, high)

		for (*nums)[high] >= temp {
			high--
		}
		swap(nums, low, high)
	}
	return low
}

func ShellSort(nums *[]int) {
	increment := len(*nums) - 1
	gap := 3
	for increment > 1 {
		increment = increment/gap + 1
		for i := 0; i < len(*nums); i += increment {
			j := i
			for j > 0 {
				if (*nums)[j-increment] > (*nums)[j] {
					// if L.r[j-increment] > L.r[j] {
					// swap(L, j-increment, j)
					swap(nums, j-increment, j)
				}
				j -= increment
			}
		}
	}
}

// heap sort未懂

func HeapSort(nums *[]int) {
	tosort := (*nums)
	buildMaxHeap(tosort)
	for i := len(tosort) - 1; i >= 0; i-- {
		tosort[i], tosort[0] = tosort[0], tosort[i]
		maxHeapify(tosort[:i], 0)
	}
}

func buildMaxHeap(tosort []int) {
	// from http://en.wikipedia.org/wiki/Heapsort
	// iParent = floor((i-1) / 2)
	for i := (len(tosort) - 1) / 2; i >= 0; i-- {
		maxHeapify(tosort, i)
	}
}
func maxHeapify(tosort []int, position int) {
	size := len(tosort)
	maximum := position
	leftChild := 2*position + 1
	rightChild := leftChild + 1
	if leftChild < size && tosort[leftChild] > tosort[position] {
		maximum = leftChild
	}
	if rightChild < size && tosort[rightChild] > tosort[maximum] {
		maximum = rightChild
	}

	if position != maximum {
		tosort[position], tosort[maximum] = tosort[maximum], tosort[position]
		maxHeapify(tosort, maximum) //recursive
	}
}

func HeapSortV2(nums *[]int) {
	array := (*nums)
	ep := (len(array) - 1) >> 1
	// fmt.Printf("ep:%v", ep)
	for i := ep; i >= 0; i-- {
		heapt(array, i, len(array)-1)
	}

	for i := len(array) - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		heapt(array, 0, i-1)
	}
}

func heapt(array []int, start int, end int) {
	le := start*2 + 1
	re := le + 1
	if le > end {
		return
	}

	var tmp = le
	if re <= end && array[re] > array[le] {
		tmp = re
	}

	if array[tmp] > array[start] {
		// fmt.Println(start, end, array)
		array[start], array[tmp] = array[tmp], array[start]
		heapt(array, tmp, end)
	}
}
