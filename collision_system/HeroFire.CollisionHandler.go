package collision_system

import (
	"fmt"
	"strings"
)

type HeroFireCollisionHandler struct {
	CollisionHandlerInterface
}

func NewHeroFireCollisionHandler(next CollisionHandlerInterface) CollisionHandlerInterface {
	herofireCollisionHdr := &HeroFireCollisionHandler{}
	collisionHdr := NewCollisionHandler(herofireCollisionHdr)
	if next != nil {
		collisionHdr.SetNext(next)
	}
	return collisionHdr
}

func (herofireCollisionHdr *HeroFireCollisionHandler) Collision(c1 SpriteInterface, c2 SpriteInterface, world *World) {
	isC1Hero := strings.Compare(c1.Display(), "H") == 0
	var fire *Fire
	var hero *Hero
	if isC1Hero {
		hero = c1.(*Hero)
		fire = c2.(*Fire)
	} else {
		hero = c2.(*Hero)
		fire = c1.(*Fire)
	}
	world.Remove(fire.GetCoord())
	hero.SetHp(hero.GetHp() - 10)
	if hero.GetHp() == 0 {
		world.Remove(hero.GetCoord())
		fmt.Printf("hero %v die\n", hero)
		return
	}
	if isC1Hero {
		world.Remove(hero.GetCoord())
		newHeroCoord := fire.GetCoord()
		world.spriteMap[newHeroCoord] = hero
		hero.SetCoord(newHeroCoord)
	}
}

func (herofireCollisionHdr *HeroFireCollisionHandler) IsMatch(c1 SpriteInterface, c2 SpriteInterface) bool {
	displayType := fmt.Sprintf("%v%v", c1.Display(), c2.Display())
	return strings.Compare(displayType, "HF") == 0 || strings.Compare(displayType, "FH") == 0
}
