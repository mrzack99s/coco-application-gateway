package utils

import (
	"regexp"

	"github.com/mrzack99s/coco-application-gateway/internal/types"
	"github.com/mrzack99s/coco-application-gateway/internal/vars"
)

func FindEndpointMatchPathMatch(hostname, rpath string, https bool) *types.RouteEndpointType {

	routes := vars.HTTPRouting
	if https {
		routes = vars.HTTPSRouting
	}

	if _, ok := routes["_default"]; ok {
		for routePath := range routes["_default"] {
			newPath := "^" + routePath + "$"
			regex := regexp.MustCompile(newPath)
			match := regex.Match([]byte(rpath))
			if !match {
				continue
			} else {
				endpoint := routes["_default"][routePath]
				return &endpoint
			}
		}

	}

	for routePath := range routes[hostname] {
		newPath := "^" + routePath + "$"

		regex := regexp.MustCompile(newPath)
		match := regex.Match([]byte(rpath))
		if !match {
			continue
		} else {
			endpoint := routes[hostname][routePath]
			return &endpoint
		}
	}

	return nil
}
