package algoritmer

import "fmt"

// Les https://en.wikipedia.org/wiki/Bubble_sort
func Bubble_sort_modified(list []int) {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

func main() {
	n := []int{1, 199, 3, 2, 5, 80, 99, 500}
	fmt.Println("Før: ", n)
	Bubble_sort_modified(n)
	fmt.Println("After: ", n)
}
