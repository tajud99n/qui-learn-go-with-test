package integer

// Add takes two integers and returns the sum of them
func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Printlin(sum)
	// Output: 6
}