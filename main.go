package main

import (
	"time"

	"github.com/subsystemio/subsystem"
	"gopkg.in/gin-gonic/gin.v1"
)

var (
	subsystems map[string]SubSystem.SubSystem
)

//StartCheck checks for SubSystems that haven't reported in for 10 seconds.
func StartCheck() chan bool {
	stop := make(chan bool)

	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				for _, sub := range subsystems {
					//Check if LastCheck was more than 10 seconds ago.
					if err := sub.HealthCheck(); err != nil {
						delete(subsystems, sub.Data.Name)
					}
				}
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()

	return stop
}

func main() {
	subsystems = make(map[string]SubSystem.SubSystem)
	StartCheck()

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/deploy", DeployPOST)
		v1.POST("/subsystems", SubSystemPOST)
	}

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
