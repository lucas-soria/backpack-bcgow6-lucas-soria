package ordenamiento

func BubbleSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}

func InsertionSort(slice []int) []int {
	var n = len(slice)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
			}
			j = j - 1
		}
	}
	return slice
}

func SelectionSort(slice []int) []int {
	var n = len(slice)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
	}
	return slice
}
