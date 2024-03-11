package main

import (
	"fmt"
)

type Hitbox interface {
//	Render() error
	CloseTo(other Hitbox) bool
	Touching(other Hitbox) bool
	ResolveCollision(other Hitbox) bool
}

type CircleHitbox struct {
	// Implements Hitbox
	X      float64
	Y      float64
	R      float64
}

func (c CircleHitbox) ResolveCollision(other Hitbox) bool {
	// TODO: Implement
	return false
	// switch other := other.(type) {
	// 	case CircleHitbox:
	// 		dx := c.X - other.X
	// 		fmt.Println(dx)
	// 		return false
	// 	default:
	// 		// TODO: Implement other hitbox types
	// 		fmt.Println("Hitbox type not implemented")
	// 		return false
	// }
	// return false
}

func (c CircleHitbox) CloseTo(other Hitbox) bool {
	// TODO: Implement for performance
	return true
	// switch other := other.(type) {
	// 	case CircleHitbox:
	// 		return true
	// 	default:
	// 		// TODO: Implement other hitbox types
	// 		fmt.Println("Hitbox type not implemented")
	// 		return false
	// }
	// return false
}

func (c CircleHitbox) Touching(other Hitbox) bool {
	if !c.CloseTo(other) {
		return false
	}
	switch other := other.(type) {
		case *CircleHitbox:
			dx := c.X - other.X
			dy := c.Y - other.Y
			dr := c.R + other.R
			return dx*dx+dy*dy <= dr*dr
		default:
			// TODO: Implement other hitbox types
			fmt.Println("Hitbox type not implemented")
			return false
	}
	return false
}