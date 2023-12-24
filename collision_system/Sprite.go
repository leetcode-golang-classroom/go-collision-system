package collision_system

import "log"

type Sprite struct {
	SpriteInterface
	coord int
}
type SpriteInterface interface {
	Display() string
	GetCoord() int
	SetCoord(coord int)
}

func NewSprite() *Sprite {
	return &Sprite{}
}
func (sprite *Sprite) SetCoord(coord int) {
	if coord < 0 || coord > 29 {
		log.Fatalf("%v is not in range [0, 29]", coord)
	}
	sprite.coord = coord
}

func (sprite *Sprite) GetCoord() int {
	return sprite.coord
}
