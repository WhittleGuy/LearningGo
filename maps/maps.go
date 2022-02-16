package maps

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }

type Dictionary map[string]string // Never initialize empty or runtime panic

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryErr("Could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("Cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("Cannot update word because it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition // Do not need pointer to modify maps
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
