package main

import (
	"log"
	"os/exec"

	"gopkg.in/gin-gonic/gin.v1"
)

func runCommand(args []string) {
	cmd := exec.Command("cmd", args...)
	cmd.Run()
}

func (s *Server) DeployPOST(c *gin.Context) {
	runCommand([]string{"/k", "go", "get", "github.com/subsystemio/sub-hello"})
	go runCommand([]string{"/k", "sub-hello"})

	log.Printf("Started %v\n", "sub-hello")
}

func SubSystemPOST(c *gin.Context) {
	var d SubSystem.SubSystemData
	if c.BindJSON(&d) == nil {

		s := SubSystem.SubSystem{
			Data: d,
			Port: 9000 + len(subsystems),
		}

		subsystems[d.Name] = s
		c.String(200, string(s.Port))
	} else {
		c.JSON(500, gin.H{
			"error": "Failed to load SubSystem data.",
		})
	}
}
