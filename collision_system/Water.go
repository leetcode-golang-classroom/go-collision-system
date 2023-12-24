package collision_system

import "fmt"

type Water struct {
	*Sprite
}

func NewWater(coord int) *Water {
	water := &Water{
		NewSprite(),
	}
	water.SetCoord(coord)
	return water
}
func (water *Water) String() string {
	return fmt.Sprintf("[%v:%v]", water.Display(), water.coord)
}
func (water *Water) Display() string {
	return "W"
}
