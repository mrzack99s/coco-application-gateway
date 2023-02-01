package vars

import (
	"github.com/corazawaf/coraza/v2"
	"github.com/mrzack99s/coco-application-gateway/internal/types"
)

var (
	Conf types.ConfigType
	WAF  *coraza.Waf
)
