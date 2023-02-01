package features

import (
	"github.com/mrzack99s/coco-application-gateway/internal/types"
	"github.com/mrzack99s/coco-application-gateway/internal/utils"
)

var (
	IPWhiteListHttp  = make(map[string][]types.IPWhiteList)
	IPWhiteListHttps = make(map[string][]types.IPWhiteList)
)

func CheckWhiteList(hostname, ip string, https bool) bool {
	ipWhiteList := IPWhiteListHttp
	if https {
		ipWhiteList = IPWhiteListHttps
	}

	if v, ok := ipWhiteList[hostname]; ok {
		if len(v) == 0 {
			return true
		}

		for _, d := range v {
			if utils.IpInCidr(d.CIDR, ip) {
				return true
			}
		}

	}

	return false
}
