package vars

import "github.com/mrzack99s/coco-application-gateway/internal/types"

var (
	BackendPools = make(map[string]types.BackendPool)
)
