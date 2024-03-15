package main

import "testing"

func TestFactorial(t *testing.T) {
	got := recursiveFactorial(0)
	var want int64 = 1
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func TestFactorial2(t *testing.T) {
	got := recursiveFactorial(1)
	var want int64 = 1
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
func TestFactorial3(t *testing.T) {
	got := recursiveFactorial(5)
	var want int64 = 120
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
func TestFactorial4(t *testing.T) {
	got := recursiveFactorial(-20)
	var want int64 = -1
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func TestIterativeFactorial(t *testing.T) {
	got := iterativeFactorial(0)
	var want int64 = 1
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func TestIterativeFactorial2(t *testing.T) {
	got := iterativeFactorial(1)
	var want int64 = 1
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
func TestIterativeFactorial3(t *testing.T) {
	got := iterativeFactorial(5)
	var want int64 = 120
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
func TestIterativeFactorial4(t *testing.T) {
	got := iterativeFactorial(-20)
	var want int64 = -1
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func TesteAproximation(t *testing.T) {
	got := eAproximation(0)
	var want float64 = 2
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}
}
func TesteAproximation2(t *testing.T) {
	got := eAproximation(2)
	var want float64 = 2 + float64(iterativeFactorial(2))
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}
}
func TesteAproximation3(t *testing.T) {
	got := eAproximation(3)
	var want float64 = 2 + float64(iterativeFactorial(3))
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}
}
