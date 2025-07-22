package algos

import (
	"cmp"
)

// BubbleSort выполняет сортировку пузырьком.
// Не рекомендуется для использования на больших массивах из-за O(n²) сложности.
//
// Работает с типами, поддерживающими сравнение: int, float64, string и т.д.

func BubbleSort[T cmp.Ordered](slice []T) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		isSorted := true
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
}

func SelectionSort[T cmp.Ordered](slice []T) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
	}
}
