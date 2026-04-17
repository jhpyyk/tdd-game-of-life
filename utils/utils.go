package utils

import (
	"slices"
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
	if len(mat) < 2 {
		return mat
	}
	if len(mat[0]) < 2 {
		return mat
	}
	topRowEmpty := true
	for _, col := range mat[0] {
		if col != 0 {
			topRowEmpty = false
			break
		}
	}
	if topRowEmpty {
		mat = removeTopRow(mat)
	}

	bottomRowEmpty := true
	for _, col := range mat[len(mat)-1] {
		if col != 0 {
			bottomRowEmpty = false
			break
		}
	}
	if bottomRowEmpty {
		mat = removeBottomRow(mat)
	}

	leftColEmpty := true
	for _, row := range mat {
		if row[0] != 0 {
			leftColEmpty = false
			break
		}
	}

	if leftColEmpty {
		mat = removeLeftColumn(mat)
	}

	rightColEmpty := true
	for _, row := range mat {
		if row[len(row)-1] != 0 {
			rightColEmpty = false
			break
		}
	}
	if rightColEmpty {
		mat = removeRightColumn(mat)
	}

	if topRowEmpty || bottomRowEmpty || leftColEmpty || rightColEmpty {
		mat = TrimPadding(mat)
	}
	return mat
}

func removeTopRow(mat [][]int) [][]int {
	return mat[1:]
}

func removeBottomRow(mat [][]int) [][]int {
	return mat[:len(mat)-1]
}

func removeLeftColumn(mat [][]int) [][]int {
	result := [][]int{}
	for _, row := range mat {
		result = append(result, row[1:])
	}
	return result
}

func removeRightColumn(mat [][]int) [][]int {
	result := [][]int{}
	for _, row := range mat {
		result = append(result, row[:len(row)-1])
	}
	return result
}

func RemoveTrailingZeros(arr []int) []int {
	for i, element := range slices.Backward(arr) {
		if element != 0 {
			return arr[:i+1]
		}
	}
	return []int{}
}
