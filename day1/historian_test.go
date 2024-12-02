package day1

import "testing"

func TestHistorian(t *testing.T) {
	t.Run("REAL INPUT: Historian function", func(t *testing.T) {
		got := Historian("input")
		want := 3569916

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("Historian function", func(t *testing.T) {
		got := Historian("test_input")
		want := 11

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("test get lists function", func(t *testing.T) {
		got := getLists([]string{"4 200", "1 2", "12 4", "3 200"})
		want := [2][]int{{4, 1, 12, 3}, {200, 2, 4, 200}}

		if len(got) != len(want) {
			t.Errorf("got %v want %v", got, want)
			return
		}

		for i := range want {
			for j := range want[i] {
				if got[i][j] != want[i][j] {
					t.Errorf("got %v want %v", got, want)
					return
				}
			}
		}
	})
	t.Run("test sort lists function", func(t *testing.T) {
		got := sortLists([2][]int{{4, 1, 12, 3}, {200, 2, 4, 200}})
		want := [2][]int{{1, 3, 4, 12}, {2, 4, 200, 200}}

		if len(got) != len(want) {
			t.Errorf("got %v want %v", got, want)
			return
		}

		for i := range want {
			for j := range want[i] {
				if got[i][j] != want[i][j] {
					t.Errorf("got %v want %v", got, want)
					return
				}
			}
		}
	})
}

func TestHistorianPart2(t *testing.T) {
	t.Run("REAL INPUT: Historian function", func(t *testing.T) {
		got := HistorianPart2("input")
		want := 26407426

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("Historian PART 2 function", func(t *testing.T) {
		got := HistorianPart2("test_input")
		want := 31

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("test similarity list function", func(t *testing.T) {
		got := getSimilarityLists([2][]int{{4, 200, 4, 3}, {200, 2, 4, 200}})
		want := []int{4, 400, 4, 0}

		if len(got) != len(want) {
			t.Errorf("got %v want %v", got, want)
			return
		}

		for i := range want {
			if got[i] != want[i] {
				t.Errorf("got %v want %v", got, want)
				return
			}
		}
	})

}
