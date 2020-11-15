package main

// Dictionary is a wrapper over the Search map
type Dictionary map[string]string

// ErrNotFound sets error message if the word were not found in dictionary
const (
	ErrNotFound      = DictionaryErr("could not find the word")
	ErrWordExists    = DictionaryErr("this word exists, can't add it")
	ErrWordNotExists = DictionaryErr("this word doe not exist, cant' update")
)

// DictionaryErr stores error messages
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search method to return word from the map
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add method for adding keys and words to dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update method change the definition
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordNotExists
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete method remove the definiton
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
