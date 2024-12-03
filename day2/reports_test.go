package day2

import "testing"

func TestAnalyzeReports(t *testing.T) {
	t.Run("REAL INPUT: analyze function", func(t *testing.T) {
		got := AnalyzeReports("input")
		want := 390

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("example input", func(t *testing.T) {
		got := AnalyzeReports("test_input")
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
