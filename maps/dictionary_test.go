package maps

import "testing"

func assertStrings(t *testing.T, got, want string){
	t.Helper()
	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error){
	t.Helper()
	if got != want{
		t.Errorf("got %q want %q", got.Error(), want.Error())
	}
}

func TestSearch(t *testing.T){
	dictionary := Dictionary{"test":"this is just a test"}

	t.Run("Search known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("Search unknown word", func(t *testing.T) {
		_, err := dictionary.Search("hello")
		want := ErrNotFound

		if err == nil{
			t.Fatal("Expected to get an error.")
		}

		assertError(t, err, want)
	})

}

func TestAdd(t *testing.T) {

	t.Run("Add new definitions", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		want := "this is just a test"

		got, err := dictionary.Search("test")

		if err != nil{
			t.Fatal("should find added word:", err)
		}

		if got != want{
			t.Errorf("got %q want %q", got, want)
		}
	})

}