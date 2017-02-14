package algorithms


func swap(list []int, i, j int){
	tmp := list[j]
	list[j] = list[i]
	list[i] = tmp
}

// Les https://en.wikipedia.org/wiki/Bubble_sort
//func Bubble_sort_modified(list []int){
func Bubble_sort(list []int) {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}
}

// Implementering av Bubble_sort algoritmen
func Bubble_sort_modified(list []int) {
	// find the length of list n
	swapped := true;
		for swapped {
			swapped = false
			for i := 0; i < len(list) - 1; i++ {
				if list[i + 1] < list[i] {
					swap(list, i, i + 1)
					swapped = true
			}
		}
	}
}

// Implementering av Quicksort algoritmen
func QSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}
