package main

import (
	"fmt"

	tm "github.com/buger/goterm"
	cli "github.com/fwielstra/roguetest/internal/cli"
)

type Coordinate struct {
	x int
	y int
}

type Entity struct {
	id       int
	name     string
	material byte
}

var Entities map[Coordinate]Entity

type Material int

const (
	Empty Material = iota
	Wall
)

var materialGraphics = []byte{
	Empty: '.',
	Wall:  '#',
}

var window [20][40]Material

func zeroWindow() {
	for y := range window {
		for x := range window[y] {
			window[y][x] = Empty
		}
	}
}

func drawBorder() {
	for y := range window {
		if y == 0 || y == len(window)-1 {
			for x := range window[y] {
				window[y][x] = Wall
			}
		}

		window[y][0] = Wall
		window[y][len(window[y])-1] = Wall
	}
}

func drawMaze() {
	window[9][9] = Wall
	window[9][10] = Wall
	window[9][11] = Wall
	window[9][12] = Wall
	window[9][13] = Wall
	window[9][14] = Wall
	window[10][9] = Wall
	window[11][9] = Wall
	window[11][10] = Wall
	window[11][11] = Wall
	window[11][12] = Wall
	window[11][13] = Wall
	window[11][14] = Wall
	window[11][15] = Wall
	window[11][16] = Wall
	window[10][16] = Wall
	window[9][16] = Wall
}

func initEntities() {
	Entities = make(map[Coordinate]Entity)
	Entities[Coordinate{10, 10}] = Entity{
		id:       1,
		name:     "Player",
		material: '@',
	}
}

func render() {
	tm.MoveCursor(1, 1)
	for y := range window {
		for x := range window[y] {
			rendered := materialGraphics[window[y][x]]

			if entity, ok := Entities[Coordinate{x, y}]; ok {
				rendered = entity.material
			}

			tm.Print(string(rendered))
		}

		tm.Printf("\n")
	}
	tm.Flush()
}

const Up = 'k'
const Down = 'j'
const Left = 'h'
const Right = 'l'

func findPlayer() (coordinate Coordinate, player Entity) {
	for k, v := range Entities {
		if v.name == "Player" {
			coordinate = k
			player = v
			return
		}
	}

	return
}

func canMoveTo(coord Coordinate) bool {
	return window[coord.y][coord.x] != Wall
}

func main() {
	zeroWindow()
	tm.Clear()

	drawBorder()
	drawMaze()
	initEntities()

	for {

		render()

		ascii, keyCode, err := cli.GetChar()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("ascii %d, character %s, keyCode %d\n", ascii, string(ascii), keyCode)

		// copy our entity map so we can ignore entities that don't change and update all the things
		newEntities := make(map[Coordinate]Entity)
		for k, v := range Entities {
			newEntities[k] = v
		}

		if ascii == 'q' {
			fmt.Println("bye")
			break
		}

		coordinate, player := findPlayer()
		newCoordinate := coordinate

		if ascii == Up {
			newCoordinate.y -= 1
		}

		if ascii == Down {
			newCoordinate.y += 1
		}

		if ascii == Left {
			newCoordinate.x -= 1
		}

		if ascii == Right {
			newCoordinate.x += 1
		}

		if canMoveTo(newCoordinate) {
			fmt.Printf("Moving player : %+v from %+v to %+v\n", player, coordinate, newCoordinate)

			delete(newEntities, coordinate)
			newEntities[newCoordinate] = player
		} else {
			fmt.Printf("Unable to move to %+v!11\n", newCoordinate)
		}

		Entities = newEntities
	}
}
