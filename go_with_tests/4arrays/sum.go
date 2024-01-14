package arrays

func Sum(numbers [5]int) (sum int){
	sum = 0
	for _, item := range numbers{
		sum += item;
	}
	return sum
}

