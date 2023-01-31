package parser

import (
	"os"

	"github.com/mrzack99s/coco-application-gateway/internal/types"
	"github.com/mrzack99s/coco-application-gateway/internal/vars"
	"gopkg.in/yaml.v2"
)

func ParseConfig() {
	dat, err := os.ReadFile("./conf/config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(dat, &vars.Conf)
	if err != nil {
		panic(err)
	}
}

func ParseRoutingConfig() {
	dat, err := os.ReadFile("./conf/routing.yaml")
	if err != nil {
		panic(err)
	}

	var routing types.RoutingType
	err = yaml.Unmarshal(dat, &routing)
	if err != nil {
		panic(err)
	}

	for _, http := range routing.HTTP {

		if _, ok := vars.HTTPRouting[http.Hostname]; !ok {
			vars.HTTPRouting[http.Hostname] = make(map[string]types.RouteEndpointType)
		}

		for _, route := range http.Routes {
			vars.HTTPRouting[http.Hostname][route.Path] = route
		}

	}

	for _, https := range routing.HTTPS {

		if _, ok := vars.HTTPSRouting[https.Hostname]; !ok {
			vars.HTTPSRouting[https.Hostname] = make(map[string]types.RouteEndpointType)
		}

		for _, route := range https.Routes {
			vars.HTTPSRouting[https.Hostname][route.Path] = route
		}

	}

}
