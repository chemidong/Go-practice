//Problem 13460

package main

import "fmt"

var N, M int //size of the board

var rowH, colH int //a coordinate of the hole
var m [10][10]int
var next [10][10][4][2]int
var v [9][9][9][9]int
var dir = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}
var ans int = 11

func Brute(count int, rowR int, colR int, rowB int, colB int) {
	var rR, cR, rB, cB int
	if count < 10 && v[rowR][colR][rowB][colB] == 0 {
		count++
		v[rowR][colR][rowB][colB] = 1
		for i := 0; i < 4; i++ {
			rR = next[rowR][colR][i][0]
			cR = next[rowR][colR][i][1]
			rB = next[rowB][colB][i][0]
			cB = next[rowB][colB][i][1]
			if rB == rowH && cB == colH {
				continue
			} else if rR == rowH && cR == colH {
				if ans > count {
					ans = count
				}
			} else {
				if rR == rB && cR == cB {
					switch i {
					case 0:
						if rowR > rowB {
							rR++
						} else {
							rB++
						}
					case 1:
						if rowR > rowB {
							rB--
						} else {
							rR--
						}
					case 2:
						if colR > colB {
							cR++
						} else {
							cB++
						}
					default:
						if colR > colB {
							cB--
						} else {
							cR--
						}
					}
				}
				Brute(count, rR, cR, rB, cB)
			}
		}
		v[rowR][colR][rowB][colB] = 0
	}
}

func main() {
	var str string                 //temp to scan
	var rowR, colR, rowB, colB int //initial coordiantes of two beads
	fmt.Scanf("%d %d\n", &N, &M)

	for i := 0; i < N; i++ {
		fmt.Scanln(&str)
		for j := 0; j < M; j++ {
			switch str[j] {
			case '.':
				m[i][j] = 1
			case 'O':
				m[i][j] = 2
				rowH = i
				colH = j
			case 'R':
				rowR = i
				colR = j
				m[i][j] = 1
			case 'B':
				rowB = i
				colB = j
				m[i][j] = 1
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if m[i][j] > 0 {
				for k := 0; k < 4; k++ {
					ti := i
					tj := j
					for m[ti+dir[k][0]][tj+dir[k][1]] > 0 {
						ti += dir[k][0]
						tj += dir[k][1]
						if m[ti][tj] == 2 {
							break
						}
					}
					next[i][j][k][0] = ti
					next[i][j][k][1] = tj
				}
			}
		}
	}
	Brute(0, rowR, colR, rowB, colB)
	if ans == 11 {
		ans = -1
	}
	fmt.Print(ans)
}
