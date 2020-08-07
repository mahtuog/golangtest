package maps

type DictionaryErr string

func (e DictionaryErr) Error() string{
	return string(e)
}

const(
	ErrNotFound = DictionaryErr("Could not find the word you were looking for")
	ErrWordExists = DictionaryErr("Word already exists in the dictonary")
	ErrWordDoesNotExist = DictionaryErr("Cannot update the definition for word as it does not exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error){
	definition, ok := d[word]

	if !ok{
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error{
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return  err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error{

	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = newDefinition
	default:
		return ErrWordDoesNotExist
	}

	return nil
}

func (d Dictionary) Delete(word string){
	delete(d, word)
}