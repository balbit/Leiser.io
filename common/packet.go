package main


type Packet interface {
	ToBitstring() string
}

type PlayerPacket struct {
	movement PlayerMovement
}