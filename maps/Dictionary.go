package maps

import "errors"

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

type Dictionary map[string]string

func (dictionary Dictionary) Search(key string) (string, error) {
	definition, ok := dictionary[key]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, isExist := d.Search(key)

	switch isExist {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return isExist
	}

	return nil
}
