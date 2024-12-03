package day3

import "testing"

func TestAnalyzeReports(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		got := CompileMults("test_input")
		want := 161

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("REAL INPUT: analyze function", func(t *testing.T) {
		got := CompileMults("input")
		want := 171183089

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestCompileMultsPart2(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		got := CompileMultsPart2("test_input_2")
		want := 48

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("REAL INPUT: analyze function", func(t *testing.T) {
		got := CompileMultsPart2("input")
		want := 63866497

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
