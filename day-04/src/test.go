func main_test() {
	part1Examples()
	part1Input()
	part2Examples()
	part2Input()
}

func part1Examples() {
	input := readFile("example_1")
	expected := 2
	result := solvePart1(input)
	assertEqual(expected, result)
}

func part1Input() {
	input := readFile("input")
	expected := 483
	result := solvePart1(input)
	assertEqual(expected, result)
}

func part2Examples() {
	input := readFile("example_2")
	expected := 4
	result := solvePart2(input)
	assertEqual(expected, result)
}

func part2Input() {
	input := readFile("input")
	expected := 874
	result := solvePart2(input)
	assertEqual(expected, result)
}

func assertEqual(expected, actual int) {
	if expected != actual {
		panic(fmt.Sprintf("Expected: %d, Actual: %d", expected, actual))
	}
}
