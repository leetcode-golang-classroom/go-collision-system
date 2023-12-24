package collision_system

import (
	"fmt"
	"strings"
)

type HeroWaterCollisionHandler struct {
	CollisionHandlerInterface
}

func NewHeroWaterCollisionHandler(next CollisionHandlerInterface) CollisionHandlerInterface {
	herowaterCollisionHdr := &HeroWaterCollisionHandler{}
	collsionHdr := NewCollisionHandler(herowaterCollisionHdr)
	if next != nil {
		collsionHdr.SetNext(next)
	}
	return collsionHdr
}

func (herowaterCollisionHdr *HeroWaterCollisionHandler) Collision(c1 SpriteInterface, c2 SpriteInterface, world *World) {
	isC1Hero := strings.Compare(c1.Display(), "H") == 0
	var water *Water
	var hero *Hero
	if isC1Hero {
		hero = c1.(*Hero)
		water = c2.(*Water)
	} else {
		hero = c2.(*Hero)
		water = c1.(*Water)
	}
	world.Remove(water.GetCoord())
	hero.SetHp(hero.GetHp() + 10)
	if isC1Hero {
		world.Remove(hero.GetCoord())
		newHeroCoord := water.GetCoord()
		world.spriteMap[newHeroCoord] = hero
		hero.SetCoord(newHeroCoord)
	}
}

func (herowaterCollisionHdr *HeroWaterCollisionHandler) IsMatch(c1 SpriteInterface, c2 SpriteInterface) bool {
	displayType := fmt.Sprintf("%v%v", c1.Display(), c2.Display())
	return strings.Compare(displayType, "HW") == 0 || strings.Compare(displayType, "WH") == 0
}
