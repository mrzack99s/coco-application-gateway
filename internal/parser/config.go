package parser

import (
	"fmt"
	"os"

	"github.com/mrzack99s/coco-application-gateway/internal/features"
	"github.com/mrzack99s/coco-application-gateway/internal/loadbalancer"
	"github.com/mrzack99s/coco-application-gateway/internal/types"
	"github.com/mrzack99s/coco-application-gateway/internal/vars"
	"go.uber.org/ratelimit"
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

func ParsePoolConfig() {
	dirName := "./conf/pool.d/"
	files, err := os.ReadDir(dirName)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			dat, err := os.ReadFile(fmt.Sprintf("%s%s", dirName, file.Name()))
			if err != nil {
				panic(err)
			}

			var pools types.BackendPoolConfig
			err = yaml.Unmarshal(dat, &pools)
			if err != nil {
				panic(err)
			}

			for _, v := range pools.Pools {

				vars.BackendPools[v.Name] = v
				vars.BackendPoolHealthy[v.Name] = []int{}
				loadbalancer.RR[v.Name] = &loadbalancer.LoadBalancerRR{
					BackendPoolName: v.Name,
				}
			}

			var probes types.HealthProbeConfig
			err = yaml.Unmarshal(dat, &probes)
			if err != nil {
				panic(err)
			}

			for _, v := range probes.Probes {
				vars.HealthProbe[v.Name] = v
			}

		}
	}

}

func ParseRoutingConfig() {
	dirName := "./conf/routing.d/"
	files, err := os.ReadDir(dirName)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() {

			dat, err := os.ReadFile(fmt.Sprintf("%s%s", dirName, file.Name()))
			if err != nil {
				panic(err)
			}

			var routing types.RoutingTypeConfig
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

				features.WAFHttp[http.Hostname] = http.Features.WAFEnable
				features.RateLimitHttp[http.Hostname] = http.Features.RateLimit
				if http.Features.RateLimit.RequestPerSecond > 0 {
					features.RateLimitHttp[http.Hostname] = types.RateLimit{
						RequestPerSecond: http.Features.RateLimit.RequestPerSecond,
						Limiter:          ratelimit.New(http.Features.RateLimit.RequestPerSecond),
					}
				} else {
					features.RateLimitHttp[http.Hostname] = types.RateLimit{
						RequestPerSecond: http.Features.RateLimit.RequestPerSecond,
						Limiter:          ratelimit.NewUnlimited(),
					}
				}

				features.IPWhiteListHttp[http.Hostname] = http.Features.IPWhiteList

			}

			for _, https := range routing.HTTPS {

				if _, ok := vars.HTTPSRouting[https.Hostname]; !ok {
					vars.HTTPSRouting[https.Hostname] = make(map[string]types.RouteEndpointType)
				}

				for _, route := range https.Routes {
					vars.HTTPSRouting[https.Hostname][route.Path] = route
				}

				features.WAFHttps[https.Hostname] = https.Features.WAFEnable

				features.RateLimitHttps[https.Hostname] = https.Features.RateLimit
				if https.Features.RateLimit.RequestPerSecond > 0 {
					features.RateLimitHttps[https.Hostname] = types.RateLimit{
						RequestPerSecond: https.Features.RateLimit.RequestPerSecond,
						Limiter:          ratelimit.New(https.Features.RateLimit.RequestPerSecond),
					}
				} else {
					features.RateLimitHttps[https.Hostname] = types.RateLimit{
						RequestPerSecond: https.Features.RateLimit.RequestPerSecond,
						Limiter:          ratelimit.NewUnlimited(),
					}
				}

				features.IPWhiteListHttps[https.Hostname] = https.Features.IPWhiteList

			}
		}

	}

}
