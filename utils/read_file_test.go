package utils

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	t.Run("one file line", func(t *testing.T) {
		got := ReadFile("oneline_input")
		want := []string{"000000 0"}
		assertListEquality(t, got, want)
	})
	t.Run("long file", func(t *testing.T) {
		got := ReadFile("multiline_input")
		want := []string{"000000 0", "111", "201 201"}
		assertListEquality(t, got, want)
	})
}

func assertListEquality(t testing.TB, got, want []string) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("got %q want %q", got, want)
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %q want %q", got, want)
		}
	}

}
