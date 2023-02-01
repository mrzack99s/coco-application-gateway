package types

import "go.uber.org/ratelimit"

type RateLimit struct {
	RequestPerSecond int `yaml:"requestPerSecond"`
	Limiter          ratelimit.Limiter
}

type IPWhiteList struct {
	CIDR string `yaml:"cidr"`
}
