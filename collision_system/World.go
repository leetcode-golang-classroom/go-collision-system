package collision_system

import (
	"fmt"
	"math/rand"
	"time"
)

type World struct {
	spriteMap    map[int]SpriteInterface
	collisionHdr CollisionHandlerInterface
}

func NewWord(collisionHdr CollisionHandlerInterface) *World {
	return &World{
		spriteMap:    make(map[int]SpriteInterface),
		collisionHdr: collisionHdr,
	}
}
func (world *World) Remove(coord int) {
	delete(world.spriteMap, coord)
}
func (world *World) Run() {
	world.RandomGenerate10Sprite()
	for len(world.spriteMap) != 0 {
		fmt.Printf("current map: %v\n", world.spriteMap)
		fmt.Print("please enter x1 x2:")
		var x1, x2 int
		fmt.Scanf("%d %d", &x1, &x2)
		if x2 > 29 || x2 < 0 {
			fmt.Println("x2 must in range [0, 29]")
			continue
		}
		world.Move(x1, x2)
	}
}
func (world *World) Move(x1, x2 int) {
	c1, x1exists := world.spriteMap[x1]
	c2, x2exists := world.spriteMap[x2]
	if !x1exists {
		fmt.Printf("%v dose not have sprite!", x1)
		return
	}
	if !x2exists {
		world.Remove(c1.GetCoord())
		c1.SetCoord(x2)
		world.spriteMap[x2] = c1
		fmt.Printf("%v have move to %v\n", c1, x2)
		return
	}
	fmt.Printf("c1: %v, c2: %v\n", c1, c2)
	world.collisionHdr.DoCollision(c1, c2, world)
}
func (world *World) RandomGenerate10Sprite() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for idx := 0; idx < 10; idx++ {
		fmt.Println("idx", idx)
		coord := world.RandomChooseCoord()
		world.spriteMap[coord] = world.RandomChooseSprite(coord)
	}
}

func (world *World) RandomChooseCoord() int {
	coord := 0
	coordExist := true
	for coordExist {
		coord = rand.Intn(30)
		_, coordExist = world.spriteMap[coord]
	}
	return coord
}
func (world *World) RandomChooseSprite(coord int) SpriteInterface {
	spriteType := rand.Intn(3)
	switch spriteType {
	case 0:
		return NewWater(coord)
	case 1:
		return NewFire(coord)
	default:
		return NewHero(coord)
	}
}
