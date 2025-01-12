package utils

type Index struct {
	X, Y int
}

var Directions = [...]Index{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
