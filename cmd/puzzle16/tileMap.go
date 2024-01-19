package main

type TileMap [][]rune

// Meh getters ftl
func (t TileMap) GetTile(position Position) rune {
	return t[position[0]][position[1]]
}

func (t TileMap) String() string {
	var str string

	str += "TileMap:\n"

	for _, row := range t {
		for _, tile := range row {
			str += string(tile)
		}
		str += "\n"
	}
	return str
}
