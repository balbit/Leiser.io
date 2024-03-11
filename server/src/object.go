package main

type Object interface {
//	Render() error
	Update() error
	GetHitbox() *Hitbox
}
