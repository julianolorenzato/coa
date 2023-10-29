package sort

func SelectionSort1(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				swap(&arr[i], &arr[j])
			}
		}
	}

	return arr
}

func SelectionSort2(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		swap(&arr[minIndex], &arr[i])
	}

	return arr
}

func SelectionSort3(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		if minIndex != i {
			swap(&arr[minIndex], &arr[i])
		}
	}

	return arr
}
