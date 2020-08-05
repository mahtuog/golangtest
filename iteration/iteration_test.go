package iteration

import "testing"

func TestIteration(t *testing.T) {

	assertExpectedOutput := func(t *testing.T, got, want string) {
		if got != want{
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("repeats a charater 5 times", func(t *testing.T) {
		got := Repeat("a", -10)
		want := "aaaaa"

		assertExpectedOutput(t, got, want)
	})

	t.Run("repeat n times", func(t *testing.T) {
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"

		assertExpectedOutput(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++{
		Repeat("a", 10)
	}
}