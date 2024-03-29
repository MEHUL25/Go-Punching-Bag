package sorting

import "fmt"

func BubbleSort() {
	var n = []int{1, 39, 2, 9, 7, 54, 11}

	var isDone = false

	for !isDone {
		isDone = true
		var i = 0
		for i < len(n)-1 {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
				isDone = false
			}
			i++
		}
	}

	fmt.Println(n)
}
