package integers

import "testing"

func TestAdder(t *testing.T) {
	assertCorrect := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	t.Run("adding some numbers", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		assertCorrect(t, got, want)
	})
}