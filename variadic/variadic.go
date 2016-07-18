package variadic

func Plus(nums ...int) int  { //variadic means more than one - use ...
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
