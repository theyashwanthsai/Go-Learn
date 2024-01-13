package iteration

import(

)
const c = 5
func Repeat(i string) string{
	
	var repeat string
	for i := 0; i < c; i++{
		repeat += "a"
	}
	return repeat
}