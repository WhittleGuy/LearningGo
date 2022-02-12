package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrect := func(t testing.TB, got, want int, given []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, given)
		}
	}
	t.Run("adding some numbers", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := Sum(nums)
		want := 6
		assertCorrect(t, got, want, nums)
	})

}

func TestSumAll(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("summing all", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		checkSums(t, got, want)
	})

	t.Run("summing tails", func(t *testing.T) {
		got := SumTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("summing tails with empty slices", func(t *testing.T) {
		got := SumTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
