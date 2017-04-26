package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/subsystemio/subsystem"
	"gopkg.in/gin-gonic/gin.v1"
)

func DeployPOST(c *gin.Context) {
	cmd := exec.Command("cmd", "go", "get", "-v")
	cmd.Stdin = strings.NewReader("github.com/subsystemio/subsystem")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to load SubSystem data.",
		})
		log.Fatal(err)
	}
	c.String(200, out.String())
	fmt.Printf("in all caps: %q\n", out.String())
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
