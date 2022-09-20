// Problem 1655
// max heap min heap insert() remove() find()
// parent node -> (i - 1) / 2
// left node -> i * 2 + 1 , right node = i * 2 + 2
/* Insert
step 1: add v into last position
step 2: compare v with its parent node, if needed, switch each other
step 3: repeat step 2 until no more switch needed
*/
/* Remove
step 1: remove root node
step 2: move last node into root
step 3: compare the node with childs, if needed, switch each other
step 4: repeat step 3 until no more switch needed
*/
// Find => [0] node

// med = (n_max < n_min) ? find(min) : find(max)
/* if abs(n_max - n_min) == 2
step 1: remove a node of a bigger one
step 2: insert the removed node's value to another
*/

package main

import "fmt"

const MAX bool = true
const MIN bool = false

var max [50001]int
var n_max int = -1
var min [50001]int
var n_min int = -1

var med int = -10001

func insert(v int, mode bool) {
	if mode == MAX {
		n_max++
		if n_max == 0 {
			max[0] = v
		} else {
			i := n_max
			var p int
			max[i] = v
			for i > 0 {
				p = parent(i)
				if max[p] < max[i] {
					change(&max[p], &max[i])
					i = p
				} else {
					break
				}
			}
		}
	} else {
		n_min++
		if n_min == 0 {
			min[0] = v
		} else {
			i := n_min
			var p int
			min[i] = v
			for i > 0 {
				p = parent(i)
				if min[p] > min[i] {
					change(&min[p], &min[i])
					i = p
				} else {
					break
				}
			}
		}
	}
}
func remove(mode bool) int {
	var ret int
	if mode == MAX {
		ret = max[0]
		max[0] = max[n_max]
		n_max--
		i := 0
		for {
			left, right := child(i)
			if right <= n_max {
				if max[left] > max[right] {
					if max[left] > max[i] {
						change(&max[i], &max[left])
						i = left
					} else {
						break
					}
				} else {
					if max[right] > max[i] {
						change(&max[i], &max[right])
						i = right
					} else {
						break
					}
				}
			} else if left == n_max {
				if max[left] > max[i] {
					change(&max[i], &max[left])
				}
				break
			} else {
				break
			}
		}
	} else {
		ret = min[0]
		min[0] = min[n_min]
		n_min--
		i := 0
		for {
			left, right := child(i)
			if right <= n_min {
				if min[left] < min[right] {
					if min[left] < min[i] {
						change(&min[i], &min[left])
						i = left
					} else {
						break
					}
				} else {
					if min[right] < min[i] {
						change(&min[i], &min[right])
						i = right
					} else {
						break
					}
				}
			} else if left == n_min {
				if min[left] < min[i] {
					change(&min[i], &min[left])
				}
				break
			} else {
				break
			}
		}
	}
	return ret
}

func find(mode bool) int {
	if mode == MAX {
		return max[0]
	} else {
		return min[0]
	}
}

func median() int {
	if n_max < n_min {
		return find(MIN)
	} else {
		return find(MAX)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func child(i int) (left int, right int) {
	left = i*2 + 1
	right = i*2 + 2
	return
}

func change(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

func main() {
	var N int
	var v int
	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&v)
		if med < v {
			insert(v, MIN)
		} else {
			insert(v, MAX)
		}
		if n_max > n_min && n_max-n_min == 2 {
			v = remove(MAX)
			insert(v, MIN)
		} else if n_max < n_min && n_min-n_max == 2 {
			v = remove(MIN)
			insert(v, MAX)
		}
		med = median()
		fmt.Println(med)
	}
}
