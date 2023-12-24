package collision_system

import (
	"fmt"
	"strings"
)

type FireFireCollisionHandler struct {
	CollisionHandlerInterface
}

func NewFireFireCollisionHandler(next CollisionHandlerInterface) CollisionHandlerInterface {
	firefireCollisionHdr := &FireFireCollisionHandler{}
	collsionHdr := NewCollisionHandler(firefireCollisionHdr)
	if next != nil {
		collsionHdr.SetNext(next)
	}
	return collsionHdr
}

func (firefireCollisionHdr *FireFireCollisionHandler) Collision(c1 SpriteInterface, c2 SpriteInterface, world *World) {
	fmt.Printf("Move failed for %v to %v\n", c1, c2)
}

func (firefireCollisionHdr *FireFireCollisionHandler) IsMatch(c1 SpriteInterface, c2 SpriteInterface) bool {
	return strings.Compare(fmt.Sprintf("%v%v", c1.Display(), c2.Display()), "FF") == 0
}
