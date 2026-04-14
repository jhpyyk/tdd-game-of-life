package utils

import (
	"fmt"
	"strings"
)

func StripPattern(pattern string) string {
	var sb strings.Builder
	for _, c := range pattern {
		if c == '.' || c == '#' {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

func FrobeniusInnerProduct(mat1, mat2 [][]int) int {
	result := 0
	for i := range mat1 {
		for j := range mat2 {
			result = result + mat1[i][j]*mat2[i][j]
		}
	}
	return result
}

func PadMatrix(mat [][]int) [][]int {
	padded := [][]int{}
	for range len(mat) + 2 {
		row := []int{}
		for range len(mat[0]) + 2 {
			row = append(row, 0)
		}
		padded = append(padded, row)
	}
	for i, row := range mat {
		for j := range row {
			padded[i+1][j+1] = mat[i][j]
		}
	}
	return padded
}

func ZeroMatrix(rows, cols int) [][]int {
	mat := [][]int{}
	for range rows {
		row := []int{}
		for range cols {
			row = append(row, 0)
		}
		mat = append(mat, row)
	}
	return mat
}

func GetSubMatrix(mat [][]int, x, y, sizeX, sizeY int) [][]int {
	subMat := [][]int{}

	for i := x; i < x+sizeX; i++ {
		row := []int{}
		for j := y; j < y+sizeY; j++ {
			row = append(row, mat[i][j])
		}
		subMat = append(subMat, row)
	}
	return subMat
}

func TrimPadding(mat [][]int) [][]int {
	topRowEmpty := true
	for _, col := range mat[0] {
		if col != 0 {
			topRowEmpty = false
			break
		}
	}
	if topRowEmpty {
		mat = mat[1:]
		fmt.Println(mat)
	}

	bottomRowEmpty := true
	for _, col := range mat[len(mat)-1] {
		if col != 0 {
			bottomRowEmpty = false
			break
		}
	}
	if bottomRowEmpty {
		mat = mat[:len(mat)-1]
		fmt.Println(mat)
	}

	leftColEmpty := true
	for _, row := range mat {
		if row[0] != 0 {
			leftColEmpty = false
			break
		}
	}

	result := [][]int{}
	if leftColEmpty {
		for _, row := range mat {
			result = append(result, row[1:])
		}
	} else {
		result = mat
	}

	rightColEmpty := true
	for _, row := range result {
		if row[len(row)-1] != 0 {
			rightColEmpty = false
			break
		}
	}

	result2 := [][]int{}
	if rightColEmpty {
		for _, row := range result {
			result2 = append(result2, row[:len(result)])
		}
	} else {
		result2 = result
	}

	if topRowEmpty || bottomRowEmpty || leftColEmpty || rightColEmpty {
		result2 = TrimPadding(result2)
	}

	return result2
}
