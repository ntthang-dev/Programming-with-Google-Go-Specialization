package main

type sortTest struct {
	description string
	input       []int
	output      []int
}

var sortTestCases = []sortTest{
	{
		description: "Test 5 input",
		input:       []int{3, 5, 2, 1, 4},
		output:      []int{1, 2, 3, 4, 5},
	},
}
