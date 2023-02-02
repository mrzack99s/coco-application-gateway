package handler

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/coco-application-gateway/internal/constants"
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

		path := request.URL.Path
		uri := request.URL.RequestURI()

		var endpoint types.RouteEndpointType
		var foundEndpoint bool
		wafEnable := false

		if sslEnable {
			endpoint, foundEndpoint = utils.FindEndpointMatchPathMatch(hostname, path, true)
			if isEnable, ok := features.WAFHttps[utils.FindMatchHostname(hostname, true)]; ok {
				wafEnable = isEnable
			}

			if limiter, ok := features.RateLimitHttps[utils.FindMatchHostname(hostname, true)]; ok {
				limiter.Limiter.Take()
			}

			if !features.CheckWhiteList(utils.FindMatchHostname(hostname, true), ctx.ClientIP(), true) {
				return403(ctx, fmt.Sprintf("sorry, your ip %s is not authorized", ctx.ClientIP()))
				return
			}

		} else {
			endpoint, foundEndpoint = utils.FindEndpointMatchPathMatch(hostname, path, false)
			if isEnable, ok := features.WAFHttp[utils.FindMatchHostname(hostname, false)]; ok {
				wafEnable = isEnable
			}

			if limiter, ok := features.RateLimitHttp[utils.FindMatchHostname(hostname, true)]; ok {
				limiter.Limiter.Take()
			}

			if !features.CheckWhiteList(utils.FindMatchHostname(hostname, true), ctx.ClientIP(), false) {
				return403(ctx, fmt.Sprintf("sorry, your ip %s is not authorized", ctx.ClientIP()))
				return
			}
		}

		regex := regexp.MustCompile(`/[\w-/]*`)
		matchBaseURL := regex.FindString(endpoint.Path)
		uri = strings.Replace(uri, matchBaseURL, "", 1)
		if len(uri) == 0 || uri[0] != '/' {
			uri = "/" + uri
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

		if foundEndpoint {

			switch endpoint.Action {
			case constants.ACTION_LOAD_BALANCER:
				beIndex := loadbalancer.RR[endpoint.BackendPoolName].Next()
				if beIndex == -1 {
					return502(ctx)
					return
				}

				be := vars.BackendPools[endpoint.BackendPoolName].Servers[beIndex]

				var fullUrl string
				if endpoint.Https {
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

			case constants.ACTION_PROXY:

				resp, err := utils.HttpJSONRequestWithBytesResponse(
					request.Method,
					fmt.Sprintf("%s%s", endpoint.To, uri),
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
			case constants.ACTION_REDIRECT:
				ctx.Redirect(302, endpoint.To)
			case constants.ACTION_STRING:
				ctx.String(endpoint.Response.StatusCode, endpoint.Response.Message)
			default:
				return502(ctx)
			}

		} else {
			return502(ctx)
		}

	}
}
