package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1,5)
	fmt.Println(sum)
	// Output: 6
}
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