package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// only needed below for sample processing

type Server struct {
	Connections map[string]net.Conn
}

type Message struct {
	Action string
	Data   string
}

func (s *Server) HandleConnection(conn net.Conn) {
	enc := json.NewEncoder(conn)
	dec := json.NewDecoder(conn)

	var message Message
	dec.Decode(&message)

	log.Println("First:", message.Data)
	enc.Encode(Message{Action: "Message", Data: "Reply"})
	s.Connections[message.Data] = conn

	for {
		err := dec.Decode(&message)
		if err == io.EOF {
			log.Println("Disconnected")
			conn.Close()
			return
		}

		switch message.Action {
		case "Message":
			log.Println("Message:", message.Data)
			enc.Encode(message)
		case "Close":
			conn.Close()
		}
	}
}
func (s *Server) Listen() error {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		return err
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		go s.HandleConnection(conn)
	}

}

func (s *Server) Run() {
	s.Listen()
}

func NewServer() *Server {
	server := &Server{
		Connections: make(map[string]net.Conn),
	}

	return server
}

func runCommand(args []string) {
	cmd := exec.Command("cmd", args...)
	cmd.Run()
}

func main() {

	log.Println("Launching server...")
	server := NewServer()

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/subsystems", func(c *gin.Context) {
			runCommand([]string{"/k", "go", "get", "github.com/subsystemio/sub-hello"})

			sub := new(SubSystem.SubSystem)

			go runCommand([]string{"/k", "sub-hello"})

			log.Printf("Started %v\n", "sub-hello")
		})
		v1.GET("/subsystems", func(c *gin.Context) {
			c.JSON(200, server.Connections)
		})
	}

	go server.Run()
	r.Run("localhost:8080")
}
