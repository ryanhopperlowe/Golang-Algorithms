package module02

import (
	"sort"
	"strings"
)

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	for sweep := 0; sweep < len(list); sweep++ {
		swapped := false
		for i := 0; i < len(list)-1-sweep; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSortString is a bubble sort for string slices. Try implementing it for
// practice.
func BubbleSortString(list []string) {
	for sweep := 0; sweep < len(list); sweep++ {
		swapped := false
		for i := 0; i < len(list)-1-sweep; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func (this *Person) CompareTo(p Person) int {
	if this.Age != p.Age {
		return this.Age - p.Age
	}

	if this.LastName != p.LastName {
		return strings.Compare(this.LastName, p.LastName)
	}
	return strings.Compare(this.FirstName, p.FirstName)
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {
	for sweep := 0; sweep < len(people); sweep++ {
		swapped := false
		for i := 0; i < len(people)-i-sweep; i++ {
			if people[i].CompareTo(people[i+1]) > 0 {
				people[i], people[i+1] = people[i+1], people[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
	for sweep := 0; sweep < list.Len(); sweep++ {
		swapped := false
		for i := 0; i < list.Len()-1-sweep; i++ {
			if list.Less(i+1, i) {
				list.Swap(i+1, i)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
