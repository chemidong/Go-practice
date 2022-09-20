// Problem 1655
package main

import "fmt"

var a []int

func bi(r int) {
	if r > 0 {
		v := a[r]
		low := 0
		high := r - 1
		for low < high {
			mid := (low + high) / 2
			if a[mid] > v {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		if a[low] <= v {
			low++
		}
		copy(a[low+1:], a[low:])
		a[low] = v
	}
}
func main() {
	var N int
	fmt.Scanln(&N)
	a = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&a[i])
		bi(i)
		fmt.Println(a[i/2])
	}
}
