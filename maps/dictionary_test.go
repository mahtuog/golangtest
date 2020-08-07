package maps

import (
	"testing"
)

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

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string){
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil{
		t.Fatal("should find added word:", err)
	}

	if got != definition{
		t.Errorf("got %q want %q", got, definition)
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
	dictionary := Dictionary{}


	t.Run("new word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
	})

}

func TestUpdate(t *testing.T){

	t.Run("update existing word", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "test"
		definition := "this is just a test"

		dictionary.Add(word, definition)

		newDefinition := "new definiton for test"

		dictionary.Update(word, newDefinition)

		got, _ := dictionary.Search(word)
		assertStrings(t, got, newDefinition)
	})

	t.Run("update unknown word", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "unknown word"
		newDefinition := "this word is unknown"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T){

	t.Run("deletes word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: word, definition : definition}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)

		if err != ErrNotFound{
			t.Errorf("Expected %q to be deleted", word)
		}
	})

}