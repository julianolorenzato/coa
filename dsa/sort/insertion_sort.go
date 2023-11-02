package sort

func InsertionSort1(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				swap(&arr[j], &arr[j-1])
			} else {
				break
			}
		}
	}

	return arr
}

//func InsertionSort2(arr []int) []int {
//	for i := 0; i < len(arr); i++ {
//		el := arr[i]
//
//		for j := i - 1; j > 0; j-- {
//
//		}
//	}
//}
