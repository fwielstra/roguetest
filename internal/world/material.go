package world

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
