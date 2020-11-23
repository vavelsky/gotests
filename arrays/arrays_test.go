package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("collection of two slices", func(t *testing.T) {
		got := SumAllTails([]int{0, 9}, []int{3, 7})
		want := []int{9, 7}
		checkSums(t, got, want)
	})

	t.Run("collection of empty slice and not empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}

func ExampleSumAllTails() {
	sliceOne := []int{0, 9}
	sliceTwo := []int{3, 7}
	sum := SumAllTails(sliceOne, sliceTwo)
	fmt.Println(sum)
	// Output: [9 7]
}
