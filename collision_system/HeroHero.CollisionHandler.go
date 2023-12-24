package collision_system

import (
	"fmt"
	"strings"
)

type HeroHeroCollisionHandler struct {
	// CollisionHandlerInterface
	*CollisionHandler
}

func NewHeroHeroCollisionHandler(next CollisionHandlerInterface) CollisionHandlerInterface {
	heroheroCollisionHdr := &HeroHeroCollisionHandler{}
	collisionHdr := NewCollisionHandler(heroheroCollisionHdr)
	if next != nil {
		collisionHdr.SetNext(next)
	}
	return collisionHdr
}

func (heroheroCollisionHdr *HeroHeroCollisionHandler) Collision(c1 SpriteInterface, c2 SpriteInterface, world *World) {
	fmt.Printf("Move failed for %v to %v\n", c1, c2)
}

func (heroheroCollisionHdr *HeroHeroCollisionHandler) IsMatch(c1 SpriteInterface, c2 SpriteInterface) bool {
	return strings.Compare(fmt.Sprintf("%v%v", c1.Display(), c2.Display()), "HH") == 0
}
