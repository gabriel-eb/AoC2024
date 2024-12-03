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

	t.Run("convertLinesToInts returns only ints", func(t *testing.T) {
		got := convertLinesToInts([]string{"1", "10", "999"})
		want := []int{1, 10, 999}

		if len(got) != len(want) {
			t.Errorf("got %v want %v", got, want)
		}

		for i := range want {
			if got[i] != want[i] {
				t.Errorf("got %v want %v", got, want)
				return
			}
		}
	})
}

func TestAnalyzeReportsPart2(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		got := AnalyzeReportsPart2("test_input")
		want := 10

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("REAL INPUT: analyze function", func(t *testing.T) {
		got := AnalyzeReportsPart2("input")
		want := 439

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
