package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/coco-application-gateway/internal/features"
	"github.com/mrzack99s/coco-application-gateway/internal/handler"
	"github.com/mrzack99s/coco-application-gateway/internal/health"
	"github.com/mrzack99s/coco-application-gateway/internal/logger"
	"github.com/mrzack99s/coco-application-gateway/internal/parser"
	"github.com/mrzack99s/coco-application-gateway/internal/vars"
)

func RunAppGW() {
	parser.ParseConfig()
	parser.ParsePoolConfig()
	parser.ParseRoutingConfig()

	health.ServeCheckBackendHealth()
	features.InitWAF()

	gin.DisableConsoleColor()
	logFile, _ := os.OpenFile("./logs/gin.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0655)
	wafLogFile, _ := os.OpenFile("./logs/waf.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0655)
	logger.WAFLogger = log.New(wafLogFile, "", 0)

	gin.DefaultWriter = io.MultiWriter(logFile)

	if vars.Conf.Properties.SSL.Enable {
		go func() {
			r := gin.New()
			r.Use(logger.GetGinLog(true))
			r.NoRoute(handler.Serve(true))
			r.RunTLS(fmt.Sprintf(":%d", vars.Conf.Properties.Port.Https), vars.Conf.Properties.SSL.CertPath, vars.Conf.Properties.SSL.KeyPath)
		}()
	}

	r := gin.New()
	r.Use(logger.GetGinLog(false))
	r.NoRoute(handler.Serve(false))
	r.Run(fmt.Sprintf(":%d", vars.Conf.Properties.Port.Http))
}
