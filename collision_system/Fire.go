package collision_system

import "fmt"

type Fire struct {
	*Sprite
}

func NewFire(coord int) *Fire {
	fire := &Fire{
		NewSprite(),
	}
	fire.SetCoord(coord)
	return fire
}

func (fire *Fire) String() string {
	return fmt.Sprintf("[%v:%v]", fire.Display(), fire.coord)
}
func (fire *Fire) Display() string {
	return "F"
}
