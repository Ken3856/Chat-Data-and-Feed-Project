// Author: Kenneth Johnson
// Description: Chata Data and feed project
// Date:  24/02/2024

package main // Package declaration

import (
	"fmt"      // Imports the fmt package for formated I/O
	"io"       // Imports the io package for I/O operations
	"log"      // Imports package for logging
	"net/http" // Imports package for HTTP server functionailty

	"golang.org/x/net/websocket" // Imports the web socket package
)

type Server struct { // Server struct to represent the websocket server
	conns map[*websocket.Conn]bool // maps to store the websockete connection
}

func NewServer() *Server { // This function created and returns a newserver instance
	return &Server{
		conns: make(map[*websocket.Conn]bool), // Initialises a server struct with an wmpty map
	}
}

func (s *Server) handleWS(ws *websocket.Conn) { //This method handles web socket connections
	fmt.Println("New incoming connection from client:", ws.RemoteAddr()) // Prints thre clients address.
	s.conns[ws] = true                                                   // Adds a web socket connection to the map of connections

	s.readLoop(ws) // Starts reading th messages from the client

}

func (s *Server) readLoop(ws *websocket.Conn) { // Cotinuous loop to read the messages from the web socket connnection.
	buf := make([]byte, 1024) // Buffer to store the received data.
	for {
		n, error := ws.Read(buf) // Read the data from the websocket connection
		if error != nil {        // Error handling to check for any read errors
			if error == io.EOF { // No crazy loops, if EOF, break the loop
				break
			}
			log.Println("Read error:", error) // Log any other read errors.
			break                             // break on error
		}
		msg := buf[:n]                                // Extract the error from the buffer
		fmt.Println("Received message:", string(msg)) // Prints the received message.
		// fmt.Println(string(msg))
		ws.Write([]byte("Thank you for the message!!!")) // Sends an acknowledgment me ssage to the client

		if error != nil { // Check for any write errors
			log.Println("Write error: ", error) // Log any write errors
			break                               // Break loop on error
		}
	}
}

func main() {
	server := NewServer()                                  // Create new server instance
	http.Handle("/ws", websocket.Handler(server.handleWS)) // Handle web socket connections

	// http.ListenAndServe(":3000", nil)
	error := http.ListenAndServe(":3000", nil) // Start the HTTP server on port 3000

	if error != nil { // Check for any server start errors
		log.Fatal("Error starting server: ", error) // Log and exit on error
	}

}
