package collision_system

import "fmt"

type Hero struct {
	*Sprite
	hp int
}

func NewHero(coord int) *Hero {
	hero := &Hero{
		NewSprite(),
		30,
	}
	hero.SetCoord(coord)
	return hero
}
func (hero *Hero) SetHp(hp int) {
	hero.hp = hp
}
func (hero *Hero) GetHp() int {
	return hero.hp
}
func (hero *Hero) String() string {
	return fmt.Sprintf("[%v:%v hp:%v]", hero.Display(), hero.coord, hero.hp)
}
func (hero *Hero) Display() string {
	return "H"
}
