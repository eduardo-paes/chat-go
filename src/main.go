package main

import (
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func initLogger() {
    logrus.SetFormatter(&logrus.TextFormatter{})
    logrus.SetOutput(os.Stdout)
    logrus.SetLevel(logrus.InfoLevel) // Set the desired log level (Info, Warning, Error, etc.)

	// Create the log file if doesn't exist. And append to it if it already exists.
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(logFile)
	} else {
		logrus.Error("Failed to open log file: ", err)
	}
}

var (
    upgrader = websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
    }

    // Store active WebSocket connections
    clients = make(map[*websocket.Conn]bool)
    clientMutex sync.Mutex
)

func main() {
	initLogger()

    // Initialize a Gin router
    router := gin.Default()

    logrus.Info("Server is starting...")

    // Define a route for the root endpoint
   router.GET("/ws", func(c *gin.Context) {
		// Upgrade the HTTP server connection to the WebSocket protocol
        conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
        if err != nil {
            logrus.Error("WebSocket upgrade error:", err)
            return
        }
        defer conn.Close()

		// Add the connection to the active clients list
        clientMutex.Lock()
        clients[conn] = true
        clientMutex.Unlock()

        // Handle WebSocket messages for this client
        for {
            messageType, p, err := conn.ReadMessage()
            if err != nil {
                logrus.Error("WebSocket read error:", err)

                // Remove the connection from the active clients list
                clientMutex.Lock()
                delete(clients, conn)
                clientMutex.Unlock()

                break
            }

            // Broadcast the message to all connected clients
            clientMutex.Lock()
            for client := range clients {
				if client != conn {
					if err := client.WriteMessage(messageType, p); err != nil {
						logrus.Error("WebSocket write error:", err)
						client.Close()
						delete(clients, client)
					}
				}
            }
            clientMutex.Unlock()
        }
    })

    // Run the server on port 8080
    router.Run(":8080")
}