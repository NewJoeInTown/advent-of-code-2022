package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	assert.Equal(t, solvePart1(file("example_1")), 13)
}

func TestAdjustCoord(t *testing.T) {
	assert.Equal(t, adjustCoord(Coord{1, 1}, Right), Coord{2, 1})
	assert.Equal(t, adjustCoord(Coord{1, 1}, Up), Coord{1, 0})
	assert.Equal(t, adjustCoord(Coord{1, 1}, Left), Coord{0, 1})
	assert.Equal(t, adjustCoord(Coord{1, 1}, Down), Coord{1, 2})
}

func TestMoveRope(t *testing.T) {
	assert.Equal(t,
		moveRope([]Coord{{0, 4}}, Right),
		[]Coord{{1, 4}, {0, 4}},
	)
	assert.Equal(t,
		moveRope([]Coord{{1, 4}, {0, 4}}, Right),
		[]Coord{{2, 4}, {1, 4}, {0, 4}},
	)
	assert.Equal(t,
		moveRope([]Coord{{8, 4}, {7, 4}, {6, 4}, {5, 4}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {0, 4}}, Right),
		[]Coord{{9, 4}, {8, 4}, {7, 4}, {6, 4}, {5, 4}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {0, 4}},
	)
	assert.Equal(t,
		moveRope([]Coord{{9, 4}, {8, 4}, {7, 4}, {6, 4}, {5, 4}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {0, 4}}, Right),
		[]Coord{{10, 4}, {9, 4}, {8, 4}, {7, 4}, {6, 4}, {5, 4}, {4, 4}, {3, 4}, {2, 4}, {1, 4}},
	)
}

func TestAdjustKnot(t *testing.T) {
	// Initial state
	assert.Equal(t,
		adjustKnot(Coord{0, 4}, Coord{0, 4}),
		Coord{0, 4},
	)

	// R 4 : 1
	assert.Equal(t,
		adjustKnot(Coord{1, 4}, Coord{0, 4}),
		Coord{0, 4},
	)
	// R 4 : 2
	assert.Equal(t,
		adjustKnot(Coord{2, 4}, Coord{0, 4}),
		Coord{1, 4},
	)
	// R 4 : 3
	assert.Equal(t,
		adjustKnot(Coord{3, 4}, Coord{1, 4}),
		Coord{2, 4},
	)
	// R 4 : 4
	assert.Equal(t,
		adjustKnot(Coord{4, 4}, Coord{2, 4}),
		Coord{3, 4},
	)

	// U 4 : 1
	assert.Equal(t,
		adjustKnot(Coord{4, 3}, Coord{3, 4}),
		Coord{3, 4},
	)
	// U 4 : 2
	assert.Equal(t,
		adjustKnot(Coord{4, 2}, Coord{3, 4}),
		Coord{4, 3},
	)
	// U 4 : 3
	assert.Equal(t,
		adjustKnot(Coord{4, 1}, Coord{4, 3}),
		Coord{4, 2},
	)
	// U 4 : 3
	assert.Equal(t,
		adjustKnot(Coord{4, 0}, Coord{4, 2}),
		Coord{4, 1},
	)
}

func TestMoveHeadAndTail(t *testing.T) {
	// R 4
	assert.Equal(t,
		moveHeadAndTail(Coord{0, 4}, Coord{0, 4}, &Step{Right, 4}),
		(Coord{4, 4}, Coord{3, 4}, []Coord{{0, 4}, {1, 4}, {2, 4}, {3, 4}}),
	)
	// U 4
	assert.Equal(t,
		moveHeadAndTail(Coord{4, 4}, Coord{3, 4}, &Step{Up, 4}),
		(Coord{4, 0}, Coord{4, 1}, []Coord{{3, 4}, {4, 3}, {4, 2}, {4, 1}}),
	)
}

func TestPart1Input(t *testing.T) {
	assert.Equal(t, solvePart1(file("input")), 6190)
}

func TestPart2Examples(t *testing.T) {
	assert.Equal(t, solvePart2(file("example_1")), 36)
}

func TestPart2Input(t *testing.T) {
	assert.Equal(t, solvePart2(file("input")), todo())
}
