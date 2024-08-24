package selection_sort

func SelectionSort(arr []int) []int {
	newArr := []int{}

	for range arr {
		low := searchLower(arr)
		newArr = append(newArr, arr[low])
		arr = append(arr[:low], arr[low+1:]...)
	}
	return newArr
}

func searchLower(arr []int) int {
	low := arr[0]
	lowIndice := 0

	for i := range arr {
		if arr[i] < low {
			low = arr[i]
			lowIndice = i
		}
	}

	return lowIndice
}
