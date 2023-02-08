package utils

import (
	"github.com/mrzack99s/coco-application-gateway/internal/types"
	"github.com/mrzack99s/coco-application-gateway/internal/vars"
)

func FindRuleEndpoint(hostname string, https bool) (types.RuleType, bool) {

	fhostname := FindMatchHostname(hostname, https)

	if https {
		endpoint, ok := vars.HTTPSRules[fhostname]
		return endpoint, ok
	} else {
		endpoint, ok := vars.HTTPRules[fhostname]
		return endpoint, ok
	}

}

func FindMatchHostname(hostname string, https bool) string {

	routes := vars.HTTPRules
	if https {
		routes = vars.HTTPSRules
	}

	if _, ok := routes[hostname]; ok {
		return hostname
	}

	return "_default"
}
