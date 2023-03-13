package module02

import (
	"sort"
)

// InsertionSortInt will sort a list of integers using the insertion sort
// algorithm.
//
// Big O (without binary search): O(N^2), where N is the size of the list.
// Big O (with binary search): O(N log N), where N is the size of the list.
//
// Test with: go test -run InsertionSortInt$
// The '$' at the end will ensure that the InsertionSortInterface tests won't be run.
func InsertionSortInt(list []int) {
	var sorted []int
	for _, n := range list {
		sorted = insert(sorted, n)
	}
	for i, val := range sorted {
		list[i] = val
	}
}

func insert(sorted []int, item int) []int {
	for i, n := range sorted {
		if item < n {
			after := append([]int{item}, sorted[i:]...)
			return append(sorted[:i], after...)
		}
	}
	return append(sorted, item)
}

// InsertionSortString uses insertion sort to sort string slices. Try
// implementing it for practice.
func InsertionSortString(list []string) {
	var sorted []string
	for _, str := range list {
		sorted = insertString(sorted, str)
	}
	for i, val := range sorted {
		list[i] = val
	}
}

func insertString(sorted []string, item string) []string {
	for i, sortedStr := range sorted {
		if item < sortedStr {
			after := append([]string{item}, sorted[i:]...)
			return append(sorted[:i], after...)
		}
	}
	return append(sorted, item)
}

// InsertionSortPerson uses insertion sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func InsertionSortPerson(people []Person) {
}

// InsertionSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice, but be wary that this particular algorithm is a
// little tricky to implement this way.
func InsertionSort(list sort.Interface) {
	for i := 0; i < list.Len(); i++ {
		for j := 0; j < i; j++ {
			if list.Less(i, j) {
				for k := j; k < i; k++ {
					list.Swap(i, k)
				}
			}
		}
	}

}
