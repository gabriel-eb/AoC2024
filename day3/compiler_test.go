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
	t.Run("getnumbers function", func(t *testing.T) {
		gotA, gotB := getNumbers("mul(4,123)")
		wantA, wantB := 4, 123

		if gotA != wantA && gotB != wantB {
			t.Errorf("got %d,%d want %d,%d", gotA, gotB, wantA, wantB)
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
	t.Run("getEnabledMults function", func(t *testing.T) {
		boolP := true
		got := getEnabledMults("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", &boolP)
		want := []string{"mul(2,4)", "mul(8,5)"}

		if len(got) != len(want) {
			t.Errorf("got %v want %v", got, want)
		}

		for i := range want {
			if got[i] != want[i] {
				t.Errorf("got %s want %s", got[i], want[i])
			}
		}
	})
}
