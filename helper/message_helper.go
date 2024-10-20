package helper

import (
	"log"
	"loom/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

// websocket
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan model.Message)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(c *gin.Context, db *gorm.DB) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg model.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			delete(clients, conn)
			break
		}

		msg.MessageID = uuid.New()
		broadcast <- msg

		if err = storeMessage(db, msg); err != nil {
			log.Printf("Error storing message: %v\n", err)
		}
	}
}

func HandleMessages(db *gorm.DB) {
	for msg := range broadcast {
		for client := range clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Println("Error sending message:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func storeMessage(db *gorm.DB, msg model.Message) error {
	return db.Create(&msg).Error
}
