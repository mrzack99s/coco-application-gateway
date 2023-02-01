package vars

import "github.com/mrzack99s/coco-application-gateway/internal/types"

var (
	BackendPoolHealthy = make(map[string][]int)
	HealthProbe        = make(map[string]types.HealthProbe)
)
