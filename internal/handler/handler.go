package handler

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/coco-application-gateway/internal/features"
	"github.com/mrzack99s/coco-application-gateway/internal/loadbalancer"
	"github.com/mrzack99s/coco-application-gateway/internal/types"
	"github.com/mrzack99s/coco-application-gateway/internal/utils"
	"github.com/mrzack99s/coco-application-gateway/internal/vars"
)

func Serve(sslEnable bool) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		request := ctx.Request.Clone(context.Background())
		hostname := request.Host
		if strings.Contains(hostname, ":") {
			tmp := strings.Split(hostname, ":")
			hostname = tmp[0]
		}

		uri := request.URL.RequestURI()

		var mHostname string
		var endpoint types.RuleType
		var foundEndpoint bool
		wafEnable := false

		if sslEnable {
			mHostname = utils.FindMatchHostname(hostname, true)
			if isEnable, ok := features.WAFHttps[mHostname]; ok {
				wafEnable = isEnable
			}

			if limiter, ok := features.RateLimitHttps[mHostname]; ok {
				limiter.Limiter.Take()
			}

			if !features.CheckWhiteList(mHostname, ctx.ClientIP(), true) {
				return403(ctx, fmt.Sprintf("sorry, your ip %s is not authorized", ctx.ClientIP()))
				return
			}

		} else {
			mHostname = utils.FindMatchHostname(hostname, false)
			if isEnable, ok := features.WAFHttp[mHostname]; ok {
				wafEnable = isEnable
			}

			if limiter, ok := features.RateLimitHttp[mHostname]; ok {
				limiter.Limiter.Take()
			}

			if !features.CheckWhiteList(mHostname, ctx.ClientIP(), false) {
				return403(ctx, fmt.Sprintf("sorry, your ip %s is not authorized", ctx.ClientIP()))
				return
			}
		}

		if wafEnable {
			tx := vars.WAF.NewTransaction()
			if it, err := tx.ProcessRequest(request); err != nil {
				return500(ctx, "WAF: Failed to process request")
				return
			} else if it != nil {
				auditLog := tx.AuditLog()
				returnWAFforbidden(ctx, uri, auditLog, sslEnable)
				return
			}

		}

		endpoint, foundEndpoint = vars.HTTPRules[mHostname]

		if foundEndpoint {
			beIndex := loadbalancer.RR[endpoint.Backend.PoolName].Next()
			if beIndex == -1 {
				return502(ctx)
				return
			}

			be := vars.BackendPools[endpoint.Backend.PoolName].Servers[beIndex]

			var fullUrl string
			if endpoint.Backend.Https {
				fullUrl = fmt.Sprintf("https://%s%s", be.Hostname, uri)
			} else {
				fullUrl = fmt.Sprintf("http://%s%s", be.Hostname, uri)
			}

			resp, err := utils.HttpJSONRequestWithBytesResponse(
				request.Method,
				fullUrl,
				ctx.ClientIP(),
				request.Header,
				request.Body,
			)
			if err != nil {
				return502(ctx)
				return
			}

			buf := new(bytes.Buffer)
			_, err = buf.ReadFrom(resp.Body)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			resp.Body.Close()

			response(ctx, resp.StatusCode, resp.Header.Get("Content-Type"), buf.Bytes())
		} else {
			return502(ctx)
			return
		}

	}
}
