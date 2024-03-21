package main

import "testing"

func TestSumList(t *testing.T) {
	list := []int64{1, 2, 3, 4}
	var want int64 = 10

	got := sum(list)
	assertCorrectMessage(t, float64(got), float64(want))
}

func TestSumList2(t *testing.T) {
	t.Run("adding numbers", func(t *testing.T) {
		numbers := []int64{1, 2, 3, 4, 5}
		got := sum(numbers)
		var want int64 = 15
		assertCorrectMessage(t, float64(got), float64(want))
	})
}
func TestAverage(t *testing.T) {
	list := []int64{1, 2, 3, 4}
	got := average(list)
	var want float64 = 2.5
	assertCorrectMessage(t, got, want)
}
func TestWithoutRepeated(t *testing.T) {
	t.Run("testing without repeated", func(t *testing.T) {
		list := []int64{1, 1, 34, 65, 1, 2, 44, 1, 44, 3}
		got := withoutRepeatedElements(list)
		want := []int64{1, 34, 65, 2, 44, 3}
		for i := 0; i < len(got); i++ {
			if got[i] != want[i] {
				t.Errorf("wanted %d at index %d but got %d instead", want[i], i, got[i])
			}
		}
	})
}
func assertCorrectMessage(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestWithoutRepeated2(t *testing.T) {
	t.Run("testing single repeated element slice", func(t *testing.T) {
		list := []int64{1, 1, 1, 1, 1, 1}
		got := withoutRepeatedElements(list)
		want := []int64{1}
		if len(got) != len(want) {
			t.Fatalf("error in length")

		}
		for var i := 0; i < len(want); i++{
			if got[i] != want[i] {
				t.Errorf("mismatch, expected %d at %d but got %d", got[i], i, want[i])
			}
		}
	})
}
