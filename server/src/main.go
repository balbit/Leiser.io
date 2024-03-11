package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"fmt"
	"strings"
)

// type IDType uint32

type Client struct {
    ID   IDType
    Conn *websocket.Conn       
}

// var clients = make(map[IDType]*Client) // Map of clients by ID

func onClose(clientID IDType) {
	fmt.Printf("[ID %d] Socket closed!\n", clientID);
	// DelPlayer(clientID)
}

func wsHandler(c *websocket.Conn) {
	// Generate a unique ID for the client
	clientID := IDType(uuid.New().ID())
	fmt.Printf("[ID %d] New connection\n", clientID)

    for {
        _, msg, err := c.ReadMessage()

		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
                onClose(clientID)
            } else {
				fmt.Printf("[ID %d] unknown err: %v", clientID, err)
			}
			return
		} else {

			fmt.Printf("[ID %d] msg: %s\n", clientID, string(msg))

			msg_tokens := strings.Split(string(msg), " ")
			switch msg_tokens[0] {
				case "init":
					fmt.Printf("[ID %d] good init!\n", clientID)
					c.WriteMessage(websocket.TextMessage, []byte("start_pos 50 50"))
					break
				case "keys":
					// TODO: Find better way to encode packets
					PlayerPacket := PlayerPacket{
						movement: PlayerMovement{
							up: msg_tokens[1] == "1",
							down: msg_tokens[2] == "1",
							left: msg_tokens[3] == "1",
							right: msg_tokens[4] == "1",
						}
					}
			}
		}
    }
}



func main() {
	fmt.Println("Starting server")

	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(wsHandler))

	app.Static("/", "client/public")

	app.Listen(":3000")

}
