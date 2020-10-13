package maps

import (
	"errors"
)

var errorMesssage = errors.New("Could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error){
	definition, ok := d[word]
	if !ok {
		return "", errorMesssage
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string)  {
	d[key] = value
}
