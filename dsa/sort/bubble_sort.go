package sort

func BubbleSort1(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}

	return arr
}

func BubbleSort2(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}

	return arr
}

func BubbleSort3(arr []int) []int {
	isSorted := true

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				isSorted = false
				swap(&arr[j], &arr[j+1])
			}
		}

		if isSorted {
			break
		}
	}

	return arr
}

func BubbleSort4(arr []int) []int {
	isSorted := true

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				isSorted = false
				swap(&arr[j], &arr[j+1])
			}
		}

		if isSorted {
			break
		}
	}

	return arr
}
