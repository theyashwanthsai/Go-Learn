package maps

import "errors"


type Dictionary map[string]string

var (
	NotFoundError = errors.New("Could not find the word you were looking for!")
	AlreadyExistsError = errors.New("cannot add word because it already exists")
)


func (dict Dictionary) Search(key string) (string, error){
	res, ok := dict[key]
	if !ok {
		return "", NotFoundError
	}
	return res, nil
}



func (dict Dictionary) Add(key, value string) (error){
	_, ok := dict.Search(key)
	if ok{
		return AlreadyExistsError
	}
	dict[key] = value
	return nil
}

