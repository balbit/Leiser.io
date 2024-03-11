package main

import (
	"time"
)

type PlayerMovement struct {
	up bool
	down bool
	left bool
	right bool
}

type PlayerPosition struct {
	x float64
	y float64
	r float64
}

type Player struct {
	// Implements Object
	Name string

	Hitbox *CircleHitbox
	Pos *PlayerPosition
	Movement *PlayerMovement

	LastUpdated MSTimeType
	LastPacketTime MSTimeType

}

func (p *Player) HandlePacket(packet *PlayerPacket) error {
	p.LastPacketTime = MSTimeType(time.Now().UnixNano() / int64(time.Millisecond))
	p.Movement = packet.movement
	return nil
}


func (p *Player) Update() error {
	timeDelta := float64(MSTimeType(time.Now().UnixNano() / int64(time.Millisecond)) - p.LastUpdated)
	frameCnt := timeDelta / MS_PER_TICK
	p.LastUpdated += frameCnt * MS_PER_TICK

	if frameCnt > MAX_TICKS_PER_UPDATE {
		frameCnt = MAX_TICKS_PER_UPDATE
	}

	// Update player position
	vx, vy := 0.0, 0.0
	if p.Movement.up {
		vy -= 1
	}
	if p.Movement.down {
		vy += 1
	}
	if p.Movement.left {
		vx -= 1
	}
	if p.Movement.right {
		vx += 1
	}
	if vx != 0 || vy != 0 {
		const INV_SQRT_2 = 0.70710678118
		vx *= INV_SQRT_2
		vy *= INV_SQRT_2
	}

	movementSpeed := 5.0
	movementScale := movementSpeed * frameCnt
	
	p.Pos.x += vx * movementScale
	p.Pos.y += vy * movementScale
	return nil
}

func (p *Player) GetHitbox() Hitbox {
	return p.Hitbox
}