package features

import "github.com/mrzack99s/coco-application-gateway/internal/types"

var (
	RateLimitHttp  = make(map[string]types.RateLimit)
	RateLimitHttps = make(map[string]types.RateLimit)
)
