package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Say 'aaaaa'", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertEqual(t, expected, repeated)
	})

	t.Run("Repeat character 10 times'", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		assertEqual(t, expected, repeated)
	})
}

func assertEqual(t testing.TB, expected, got string) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	repeatedChar := Repeat("z", 12)
	fmt.Println(repeatedChar)
}
