package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeat := Repeat("a")
	fmt.Println(repeat)
	// Output: aaaaa
}

func TestIter(t *testing.T) {
	assertCorrect := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("repeating some things", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"
		assertCorrect(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
