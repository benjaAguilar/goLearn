package maps

type DictionaryErr string
type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	if err == nil {
		return ErrWordExists
	}

	d[key] = value
	return nil
}

func (d Dictionary) Update(key, newVal string) error {
	_, err := d.Search(key)

	if err != nil {
		return ErrWordDoesNotExist
	}

	d[key] = newVal
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	if err != nil {
		return ErrWordDoesNotExist
	}

	delete(d, key)
	return nil
}
