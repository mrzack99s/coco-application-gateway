package vars

import "github.com/mrzack99s/coco-application-gateway/internal/types"

var (
	HTTPRouting  = make(map[string]map[string]types.RouteEndpointType)
	HTTPSRouting = make(map[string]map[string]types.RouteEndpointType)
)
