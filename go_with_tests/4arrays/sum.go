package arrays


func Sum(numbers []int) (sum int){
	sum = 0
	for _, item := range numbers{
		sum += item;
	}
	return sum
}

func SumAll(numbers ...[]int) []int{
	// lenghtOfNumbers := len(numbers)
	// sums := make([]int, lenghtOfNumbers)

	// for i, item := range numbers{
	// 	sums[i] = Sum(item)
	// }
	var sums []int
	for _, item := range numbers{
		sums = append(sums, Sum(item))
	}
	return sums
}


func SumAllTails(numbers ...[]int) []int{
	var sums []int
	for _, item := range numbers{
		// sums = append(sums, Sum(item) - item[0])
		if len(item) == 0 {
			sums = append(sums, 0)
		} else{
			tail := item[1:]
			sums = append(sums, Sum(tail))
		}
		
	}
	return sums
}
