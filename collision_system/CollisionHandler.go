package collision_system

type CollisionHandler struct {
	CollisionHandlerInterface
	Next CollisionHandlerInterface
}

type CollisionHandlerInterface interface {
	DoCollision(c1 SpriteInterface, c2 SpriteInterface, world *World)
	Collision(c1 SpriteInterface, c2 SpriteInterface, world *World)
	IsMatch(c1 SpriteInterface, c2 SpriteInterface) bool
	SetNext(next CollisionHandlerInterface)
}

func NewCollisionHandler(hdr CollisionHandlerInterface) *CollisionHandler {
	return &CollisionHandler{
		hdr,
		nil,
	}
}
func (collision *CollisionHandler) SetNext(next CollisionHandlerInterface) {
	collision.Next = next
}
func (collision *CollisionHandler) DoCollision(c1 SpriteInterface, c2 SpriteInterface, world *World) {
	collisionHdr := collision
	for collisionHdr != nil {
		if collisionHdr.IsMatch(c1, c2) {
			collisionHdr.Collision(c1, c2, world)
			break
		} else {
			collisionHdr = collisionHdr.Next.(*CollisionHandler)
		}
	}
}
