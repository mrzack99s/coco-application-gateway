package vars

import "github.com/mrzack99s/coco-application-gateway/internal/types"

var (
	HTTPRules  = make(map[string]types.RuleType)
	HTTPSRules = make(map[string]types.RuleType)
)
