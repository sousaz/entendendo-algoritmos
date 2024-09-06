package quicksort

func Quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	sorted := []int{}
	sorted = append(sorted, Quicksort(left(arr, pivot))...)
	sorted = append(sorted, pivot)
	sorted = append(sorted, Quicksort(right(arr, pivot))...)
	return sorted
}

func left(arr []int, pivot int) []int {
	newArr := []int{}
	for i := range arr {
		if arr[i] < pivot {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

func right(arr []int, pivot int) []int {
	newArr := []int{}
	for i := range arr {
		if arr[i] > pivot {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

func Quicksort2(arr []int, left int, right int) []int {
	if left < right {
		pivot := partition(arr, left, right)
		Quicksort2(arr, left, pivot-1)
		Quicksort2(arr, pivot+1, right)
	}
	return arr
}

func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	j := left - 1

	for i := left; i < right; i++ {
		if arr[i] < pivot {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j+1], arr[right] = arr[right], arr[j+1]
	return j + 1
}
