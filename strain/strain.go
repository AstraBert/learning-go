package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Discard[T any](inputList []T, f func(T) bool) []T {
	discardedList := []T{}
	for _,x := range(inputList) {
		if !f(x) {
			discardedList = append(discardedList, x)
		}
	}
	return discardedList
}

func Keep[T any](inputList []T, f func(T) bool) []T {
	keptList := []T{}
	for _,x := range(inputList) {
		if f(x) {
			keptList = append(keptList, x)
		}
	}
	return keptList
}