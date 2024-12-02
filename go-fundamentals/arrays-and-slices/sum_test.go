package arrays_and_slices

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}

	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSumAll(b *testing.B) {
	x := []int{1, 2}
	y := []int{0, 9}

	for i := 0; i < b.N; i++ {
		SumAll(x, y)
	}
}

func ExampleSumAll() {
	x := []int{1, 2}
	y := []int{0, 9}

	sums := SumAll(x, y)
	fmt.Println(sums)
	// Output: [3 9]
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func BenchmarkSumAllTails(b *testing.B) {
	x := []int{1, 2}
	y := []int{0, 9}

	for i := 0; i < b.N; i++ {
		SumAllTails(x, y)
	}
}

func ExampleSumAllTails() {
	x := []int{1, 2}
	y := []int{0, 9}

	sums := SumAllTails(x, y)
	fmt.Println(sums)
	// Output: [2 9]
}
