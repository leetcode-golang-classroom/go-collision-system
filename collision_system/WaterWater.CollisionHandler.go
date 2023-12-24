package collision_system

import (
	"fmt"
	"strings"
)

type WaterWaterCollisionHandler struct {
	CollisionHandlerInterface
}

func NewWaterWaterCollisionHandler(next CollisionHandlerInterface) CollisionHandlerInterface {
	waterwaterCollisionHdr := &WaterWaterCollisionHandler{}
	collisionHdr := NewCollisionHandler(waterwaterCollisionHdr)
	if next != nil {
		collisionHdr.SetNext(next)
	}
	return collisionHdr
}

func (waterwaterCollisionHdr *WaterWaterCollisionHandler) Collision(c1 SpriteInterface, c2 SpriteInterface, world *World) {
	fmt.Printf("Move failed for %v to %v\n", c1, c2)
}

func (waterwaterCollisionHdr *WaterWaterCollisionHandler) IsMatch(c1 SpriteInterface, c2 SpriteInterface) bool {
	return strings.Compare(fmt.Sprintf("%v%v", c1.Display(), c2.Display()), "WW") == 0
}
