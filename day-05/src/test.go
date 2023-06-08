func main_test() {
	part1Examples()
	part1Input()
	part2Examples()
	part2Input()
}

func part1Examples() {
	input := readFile("example_1")
	expected := "CMZ"
	result := solvePart1(input)
	assertEqual(expected, result)
}

func part1Input() {
	input := readFile("input")
	expected := "LJSVLTWQM"
	result := solvePart1(input)
	assertEqual(expected, result)
}

func part2Examples() {
	input := readFile("example_2")
	expected := "MCD"
	result := solvePart2(input)
	assertEqual(expected, result)
}

func part2Input() {
	input := readFile("input")
	expected := "BRQWDBBJM"
	result := solvePart2(input)
	assertEqual(expected, result)
}

func assertEqual(expected, actual string) {
	if expected != actual {
		panic(fmt.Sprintf("Expected: %s, Actual: %s", expected, actual))
	}
}
