package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/coco-application-gateway/internal/handler"
	"github.com/mrzack99s/coco-application-gateway/internal/parser"
	"github.com/mrzack99s/coco-application-gateway/internal/vars"
)

func RunAppGW() {
	parser.ParseConfig()
	parser.ParseRoutingConfig()

	if vars.Conf.Properties.SSL.Enable {
		go func() {
			r := gin.Default()
			r.NoRoute(handler.Serve(true))
			r.RunTLS(fmt.Sprintf(":%d", vars.Conf.Properties.Port.Https), vars.Conf.Properties.SSL.CertPath, vars.Conf.Properties.SSL.KeyPath)
		}()
	}

	r := gin.Default()
	r.NoRoute(handler.Serve(false))
	r.Run(fmt.Sprintf(":%d", vars.Conf.Properties.Port.Http))
}
