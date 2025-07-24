package algos

import (
	"cmp"
	"reflect"
	"testing"
)

type sortTestCase[T cmp.Ordered] struct {
	name  string
	input []T
	want  []T
}

func TestBubbleSort(t *testing.T) {
	t.Run("Ints", func(t *testing.T) {
		testSortAlgorithm(t, BubbleSort[int], getIntTestCases())
	})

	t.Run("Float", func(t *testing.T) {
		testSortAlgorithm(t, BubbleSort[float64], getFloatTestCases())
	})

	t.Run("String", func(t *testing.T) {
		testSortAlgorithm(t, BubbleSort[string], getStringTestCases())
	})
}

func TestSelectionSort(t *testing.T) {
	t.Run("Ints", func(t *testing.T) {
		testSortAlgorithm(t, SelectionSort[int], getIntTestCases())
	})

	t.Run("Float", func(t *testing.T) {
		testSortAlgorithm(t, SelectionSort[float64], getFloatTestCases())
	})

	t.Run("String", func(t *testing.T) {
		testSortAlgorithm(t, SelectionSort[string], getStringTestCases())
	})
}

func TestInsertionSort(t *testing.T) {
	t.Run("Ints", func(t *testing.T) {
		testSortAlgorithm(t, InsertionSort[int], getIntTestCases())
	})

	t.Run("Float", func(t *testing.T) {
		testSortAlgorithm(t, InsertionSort[float64], getFloatTestCases())
	})

	t.Run("String", func(t *testing.T) {
		testSortAlgorithm(t, InsertionSort[string], getStringTestCases())
	})
}

func testSortAlgorithm[T cmp.Ordered](
	t *testing.T,
	sortFunc func([]T),
	cases []sortTestCase[T],
) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]T, len(tt.input))
			copy(input, tt.input)
			sortFunc(input)
			if !reflect.DeepEqual(input, tt.want) {
				t.Errorf("got %v, want %v", input, tt.want)
			}
		})
	}
}

func getIntTestCases() []sortTestCase[int] {
	return []sortTestCase[int]{
		{"sorted ints", []int{5, 3, 6, 2, 1}, []int{1, 2, 3, 5, 6}},
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"duplicates", []int{2, 2, 1, 1}, []int{1, 1, 2, 2}},
	}
}

func getFloatTestCases() []sortTestCase[float64] {
	return []sortTestCase[float64]{
		{"sorted floats", []float64{21.53, 12.20, 6.23}, []float64{6.23, 12.20, 21.53}},
		{"empty slice", []float64{}, []float64{}},
		{"single element", []float64{.7}, []float64{.7}},
		{"duplicates", []float64{21.53, 12.20, 12.20}, []float64{12.20, 12.20, 21.53}},
	}
}

func getStringTestCases() []sortTestCase[string] {
	return []sortTestCase[string]{
		{"sorted strings", []string{"zebra", "apple", "monkey"}, []string{"apple", "monkey", "zebra"}},
		{"empty slice", []string{}, []string{}},
		{"single element", []string{"a"}, []string{"a"}},
		{"duplicates", []string{"beta", "alpha", "alpha"}, []string{"alpha", "alpha", "beta"}},
	}
}
