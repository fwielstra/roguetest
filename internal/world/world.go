package world

import (
	"strings"
)

type Room struct {
	Y, X int // position of top left corner in the parent World
	Grid [][]Material
}

type World struct {
	Rooms []Room
}

func width(lines []string) (width int) {
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}
	return width
}

func roomStringToRoom(str string) Room {
	lines := strings.Split(strings.TrimSpace(str), "\n")
	width := width(lines)
	height := len(lines)
	grid := make([][]Material, height)

	for y := range lines {
		grid[y] = make([]Material, width)
		for x := range lines[y] {
			grid[y][x] = toMaterial(lines[y][x])
		}
	}

	return Room{
		X:    0,
		Y:    0,
		Grid: grid,
	}
}

const testRoom = `
###############
#.............#
#.............#
#.............#
#.............#
#.............#
###############
`

func GenerateWorld() World {
	room := roomStringToRoom(testRoom)
	return World{
		Rooms: []Room{
			room,
		},
	}
}
