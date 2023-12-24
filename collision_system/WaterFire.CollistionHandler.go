package collision_system

import (
	"fmt"
	"strings"
)

type WaterFireCollisionHandler struct {
	CollisionHandlerInterface
}

func NewWaterFireCollisionHandler(next CollisionHandlerInterface) CollisionHandlerInterface {
	waterfireCollisionHdr := &WaterFireCollisionHandler{}
	collisionHdr := NewCollisionHandler(waterfireCollisionHdr)
	if next != nil {
		collisionHdr.SetNext(next)
	}
	return collisionHdr
}

func (waterfireCollisionHdr *WaterFireCollisionHandler) Collision(c1 SpriteInterface, c2 SpriteInterface, world *World) {
	world.Remove(c1.GetCoord())
	world.Remove(c2.GetCoord())
}

func (waterfireCollisionHdr *WaterFireCollisionHandler) IsMatch(c1 SpriteInterface, c2 SpriteInterface) bool {
	displayType := fmt.Sprintf("%v%v", c1.Display(), c2.Display())
	return strings.Compare(displayType, "WF") == 0 || strings.Compare(displayType, "FW") == 0
}
