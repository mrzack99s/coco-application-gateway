package features

import (
	"github.com/corazawaf/coraza/v2"
	"github.com/corazawaf/coraza/v2/seclang"
	"github.com/mrzack99s/coco-application-gateway/internal/vars"
)

var (
	WAFHttp  = make(map[string]bool)
	WAFHttps = make(map[string]bool)
)

func InitWAF() {

	vars.WAF = coraza.NewWaf()
	parser, _ := seclang.NewParser(vars.WAF)

	files := []string{
		"coraza.conf",
		"owasp/crs-setup.conf",
		"owasp/rules/*.conf",
	}
	for _, f := range files {
		if err := parser.FromFile(f); err != nil {
			panic(err)
		}
	}

}
