package utils

import (
	"regexp"

	"github.com/mrzack99s/coco-application-gateway/internal/types"
	"github.com/mrzack99s/coco-application-gateway/internal/vars"
)

func FindEndpointMatchPathMatch(hostname, rpath string, https bool) (types.RouteEndpointType, bool) {

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
				return endpoint, true
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
			return endpoint, true
		}
	}

	return types.RouteEndpointType{}, false
}

func FindMatchHostname(hostname string, https bool) string {

	routes := vars.HTTPRouting
	if https {
		routes = vars.HTTPSRouting
	}

	if _, ok := routes[hostname]; ok {
		return hostname
	}

	return "_default"
}
