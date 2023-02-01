package health

import (
	"fmt"
	"time"

	"github.com/mrzack99s/coco-application-gateway/internal/utils"
	"github.com/mrzack99s/coco-application-gateway/internal/vars"
)

func ServeCheckBackendHealth() {

	go func() {
		for {
			for poolKey := range vars.BackendPools {
				bePool := vars.BackendPools[poolKey]
				probe := vars.HealthProbe[bePool.HealthProbeName]

				for i, server := range bePool.Servers {
					var fullUrl string
					if probe.Https {
						fullUrl = fmt.Sprintf("https://%s%s", server.Hostname, probe.Path)
					} else {
						fullUrl = fmt.Sprintf("http://%s%s", server.Hostname, probe.Path)
					}

					resp, err := utils.HttpGETRequest(fullUrl)
					if err != nil || resp.StatusCode != probe.StatusCode {
						vars.BackendPools[poolKey].Servers[i].IsHealthy = false
						vars.BackendPoolHealthy[bePool.Name] = utils.FindAndDeleteInt(vars.BackendPoolHealthy[bePool.Name], i)
					} else {
						vars.BackendPools[poolKey].Servers[i].IsHealthy = true
						vars.BackendPoolHealthy[bePool.Name] = utils.FindAndAppendInt(vars.BackendPoolHealthy[bePool.Name], i)
					}

				}

			}
			time.Sleep(10 * time.Second)
		}
	}()

}
