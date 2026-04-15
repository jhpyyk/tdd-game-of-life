package utils_test

import (
	"slices"
	"testing"

	"github.com/jhpyyk/tdd-game-of-life/utils"
)

func TestFrobeniusInnerProduct(t *testing.T) {
	mat := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	prod := utils.FrobeniusInnerProduct(mat, mat)
	if prod != 9 {
		t.Fatalf("dot product not working, wanted %v, got %v", 9, prod)
	}
}

func TestPadMatrix(t *testing.T) {
	mat := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	want := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}

	got := utils.PadMatrix(mat)

	for i := range want {
		if !slices.Equal(got[i], want[i]) {
			t.Fatalf("Pad matrix not working, wanted %v, got %v", want, got)
		}
	}
}

func TestGetSubMatrix(t *testing.T) {
	mat := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}
	want := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	got := utils.GetSubMatrix(mat, 1, 1, 3, 3)

	for i := range want {
		if !slices.Equal(got[i], want[i]) {
			t.Fatalf("Get sub matrix not working, wanted %v, got %v", want, got)
		}
	}
}

func TestTrimPadding(t *testing.T) {
	t.Run("should trim one big block", func(t *testing.T) {
		mat := [][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 1, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 0, 0},
			{0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		}
		want := [][]int{
			{0, 1, 1, 1},
			{1, 1, 1, 1},
			{0, 1, 1, 1},
			{0, 0, 1, 0},
		}
		got := utils.TrimPadding(mat)
		for i := range want {
			if !slices.Equal(got[i], want[i]) {
				t.Fatalf("trim padding not working, wanted %v, got %v", want, got)
			}
		}
	})
	t.Run("should trim two blocks", func(t *testing.T) {
		mat := [][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 1, 1, 0, 0, 0, 0},
			{0, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 0},
			{0, 0, 0, 0, 1, 1, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		}
		want := [][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 0, 1, 1},
			{0, 0, 0, 1, 1},
		}
		got := utils.TrimPadding(mat)
		for i := range want {
			if !slices.Equal(got[i], want[i]) {
				t.Fatalf("trim padding not working, wanted %v, got %v", want, got)
			}
		}
	})
	t.Run("should trim two blocks 2", func(t *testing.T) {
		mat := [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 1, 0, 0, 1, 1, 0},
			{0, 1, 1, 0, 0, 1, 1, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		}
		want := [][]int{
			{1, 1, 0, 0, 1, 1},
			{1, 1, 0, 0, 1, 1},
		}
		got := utils.TrimPadding(mat)
		for i := range want {
			if !slices.Equal(got[i], want[i]) {
				t.Fatalf("trim padding not working, wanted %v, got %v", want, got)
			}
		}
	})
	t.Run("should trim empty", func(t *testing.T) {
		mat := [][]int{}
		want := [][]int{}
		got := utils.TrimPadding(mat)
		for i := range want {
			if !slices.Equal(got[i], want[i]) {
				t.Fatalf("trim padding not working, wanted %v, got %v", want, got)
			}
		}
	})
}
