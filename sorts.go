package algos

import (
	"cmp"
)

// BubbleSort выполняет сортировку пузырьком.
// Сложность: O(n²) в худшем случае, O(n) для уже отсортированного массива.
// Стабильный алгоритм. Не используйте для n > 1000 элементов.
// Поддерживает типы: int, float64, string и другие comparable-типы.
//
// Пример:
// arr := []int{3, 1, 2}
// BubbleSort(arr)
// fmt.Println(arr) // [1, 2, 3]

func BubbleSort[T cmp.Ordered](slice []T) {
	if len(slice) == 0 {
		return
	}
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

// SelectionSort сортирует слайс выбором.
// Сложность: O(n²) в любом случае. Минимум обменов: n-1.
// Нестабильный алгоритм. Эффективен когда обмены дорогие.
//
// Пример:
// arr := []string{"banana", "apple", "pear"}
// SelectionSort(arr)
// fmt.Println(arr) // ["apple", "banana", "pear"]

func SelectionSort[T cmp.Ordered](slice []T) {
	if len(slice) == 0 {
		return
	}
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

// InsertionSort выполняет сортировку вставками.
// Сложность: O(n²) в худшем случае, O(n) для почти отсортированных данных.
// Стабильный алгоритм. Эффективен для n < 50 или частично упорядоченных данных.
//
// Пример:
// arr := []float64{3.14, 1.0, 2.7}
// InsertionSort(arr)
// fmt.Println(arr) // [1.0, 2.7, 3.14]

func InsertionSort[T cmp.Ordered](slice []T) {
	if len(slice) <= 1 {
		return
	}
	for j := 1; j < len(slice); j++ {
		key := slice[j]
		i := j - 1
		for ; i >= 0 && slice[i] > key; i-- {
			slice[i+1] = slice[i]
		}
		slice[i+1] = key
	}
}
