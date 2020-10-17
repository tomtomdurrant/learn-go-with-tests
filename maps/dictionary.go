package maps

const (
	errorWordNotFound   = DictionaryErr("Could not find the word you were looking for")
	errWordExists = DictionaryErr("Word already exists")
	errWordDoesNotExist = DictionaryErr("could not update that word because it does not exist")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errorWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errorWordNotFound:
		d[word] = definition
	case nil:
		return errWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errorWordNotFound:
		return errWordDoesNotExist
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string)  {
	delete(d, word)
}