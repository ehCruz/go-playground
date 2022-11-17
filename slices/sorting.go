package slices

func SelectionSort(num []int) {
	arrLen := len(num)

	var min, aux int

	for i := 0; i < arrLen; i++ {
		min = i
		for j := min + 1; j < arrLen; j++ {
			if num[j] < num[min] {
				min = j
			}
		}
		if i != min {
			aux = num[i]
			num[i] = num[min]
			num[min] = aux
		}
	}
}
