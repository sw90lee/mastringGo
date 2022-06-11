package main

import (
	"math/rand"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func createMatrix(row, col int) [][]int {
	r := make([][]int, row)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			r[i] = append(r[i], random(-5, i*j))
		}
	}
	return r
}

func getCofactor(A [][]float64, temp [][]float64, p int, q int, n int) {
	i := 0
	j := 0

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if row != p && col != q {
				temp[i][j] = A[row][col]
				j++
				if j == n-1 {
					j = 0
					i++
				}
			}
		}
	}
}

func determinant(A [][]float64, n int) float64 {
	D := float64(0)
	if n == 1 {
		return A[0][0]
	}

	temp := createMatrix(n, n)
	sign := 1

	for f := 0; f < n; f++ {
		getCofactor(A, temp, 0, f, n)
		D += float64(sign) * A[0][f] * determinant(temp, n-1)
		sign = -sign
	}
	return D
}
