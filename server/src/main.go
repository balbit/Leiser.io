package main

import (
	"os"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"log"
	"fmt"
	"time"
	"math/rand"
)

type MousePosition struct {
    X int `json:"x"`
    Y int `json:"y"`
}

type Client struct {
    ID   uint32                
    Conn *websocket.Conn       
    Data map[string]interface{}
}

var loadCounter int = 0

var clients = make(map[uint32]*Client) // Map of clients by ID
var register = make(chan *Client)      // Channel to register new clients
var unregister = make(chan *Client)    // Channel to unregister clients

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func wsHandler(c *websocket.Conn) {
	// Generate a unique ID for the client
	clientID := uuid.New().ID()
	// Generate a random color for the client
	clientColor := fmt.Sprintf("#%06x", rand.Intn(0xFFFFFF))
	// Set default mouse position
	clientMousePos := MousePosition{0, 0}

	clientData := map[string]interface{}{
		"id":    clientID,
		"color": clientColor,
		"mousePos": clientMousePos,
	}
	clients[clientID] = &Client{clientID, c, clientData}
	defer delete(clients, clientID)

    for {
        _, msg, err := c.ReadMessage()
        if err != nil {
            println("read:", err)
            break
        }

		var pos MousePosition
        err = json.Unmarshal(msg, &pos)
        if err != nil {
            println("error unmarshalling message:", err)
            continue
        }

		// update the latest position
		clients[clientID].Data["mousePos"] = pos
    }
}


func startBroadcaster() {
	fmt.Println("Starting broadcaster")
    ticker := time.NewTicker(time.Millisecond * 1) // Adjust tick rate as needed
    defer ticker.Stop()

    for {
        select {
        case client := <-register:
            clients[client.ID] = client
        case client := <-unregister:
            if _, ok := clients[client.ID]; ok {
                delete(clients, client.ID)
                client.Conn.Close()
            }
        case <-ticker.C:
            broadcastUpdate()
        }
    }
}


func broadcastUpdate() {
    // Prepare the data to be sent, e.g., a list of all clients' data
    var allClientsData []map[string]interface{}
    for _, client := range clients {
        allClientsData = append(allClientsData, client.Data)
    }
    
    message, err := json.Marshal(allClientsData)
    if err != nil {
        log.Printf("Error marshalling message: %v", err)
        return
    }

    for _, client := range clients {
        if err := client.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
            log.Printf("Error sending message: %v", err)
            unregister <- client
        }
    }
}


func main() {
	fmt.Println("Starting server")
	app := fiber.New()
	go startBroadcaster()

	app.Use(func(c *fiber.Ctx) error {
		// Increment the load counter
		loadCounter++
		// Pass control to the next middleware/handler
		return c.Next()
	})

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(wsHandler))

	app.Get("/api/position", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"x": 40,
			"y": 40,
			"numLoads": loadCounter,
		})
	})
	app.Static("/", "../../client/public")

	app.Listen(getPort())

}
