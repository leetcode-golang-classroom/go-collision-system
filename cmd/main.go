package main

import (
	"github.com/leetcode-golang-classroom/go-collision-system/collision_system"
)

func main() {
	collisonHdr := collision_system.NewFireFireCollisionHandler(
		collision_system.NewHeroFireCollisionHandler(
			collision_system.NewHeroWaterCollisionHandler(
				collision_system.NewHeroHeroCollisionHandler(
					collision_system.NewWaterWaterCollisionHandler(
						collision_system.NewWaterFireCollisionHandler(
							nil,
						),
					),
				),
			),
		),
	)
	// fmt.Println(collisonHdr)
	world := collision_system.NewWord(collisonHdr)
	world.Run()
}
