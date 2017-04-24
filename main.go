package main

import (
	"log"

	"github.com/subsystemio/subsystem"
	"gopkg.in/gin-gonic/gin.v1"
)

var (
	subsystems map[string]SubSystem.SubSystem
)

func main() {
	subsystems = make(map[string]SubSystem.SubSystem)

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/subsystems", func(c *gin.Context) {
			var d SubSystem.SubSystemData
			if c.BindJSON(&d) == nil {
				log.Printf("Data %v", d)

				h := SubSystem.HealthCheck{
					Name: d.Name,
					Key:  "password",
				}
				s := SubSystem.SubSystem{
					Data:   d,
					Health: h,
				}

				subsystems[d.Name] = s

				c.JSON(200, gin.H{
					"name": s.Health.Name,
					"key":  s.Health.Key,
				})
			} else {
				c.JSON(500, gin.H{
					"error": "Failed to load SubSystem data.",
				})
			}
		})
		v1.POST("/health", func(c *gin.Context) {
			var d SubSystem.HealthCheck
			if c.BindJSON(&d) == nil {
				if val, ok := subsystems[d.Name]; ok {
					if val.Health.Key == d.Key {
						c.String(200, "Ba-Bump")
						return
					}
				}

				c.String(500, "Bad SubSystem or Key")

			} else {
				c.String(500, "Failed to bind healthcheck")
			}
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
