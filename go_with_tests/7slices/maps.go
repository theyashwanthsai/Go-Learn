package maps

import "errors"


type Dictionary map[string]string

var NotFoundError = errors.New("Could not find the word you were looking for!")

func (dict Dictionary) Search(key string) (string, error){
	res, ok := dict[key]
	if !ok {
		return "", NotFoundError
	}
	return res, nil
}