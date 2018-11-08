package world

import "fmt"

// represents a cell in a Room; used for rendering and collision detection etc.
type Material int

const (
	Empty Material = iota
	Wall
)

var MaterialGraphics = []byte{
	Empty: '.',
	Wall:  '#',
}

func toMaterial(input byte) Material {
	for material, graphic := range MaterialGraphics {
		if graphic == input {
			return Material(material)
		}
	}

	// probably fucky conversion.
	fmt.Println("Unable to find material for input byte", input)
	return Empty
}
