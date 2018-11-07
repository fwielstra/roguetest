package world

import "strings"

type Room struct {
	X, Y int // position of top left corner in the parent World
	Grid [][]Material
}

type World struct {
	Width, Height int
	Rooms         []Room
}

const firstRoom = `
#####
#...#
#...#
#...#
#####
`

func roomStringToRoom(str string) Room {
	trimmed := strings.TrimSpace(str)
	width = str
}

func GenerateWorld() World {
	room := Room{
		X: 0,
		Y: 0,
	}
}
