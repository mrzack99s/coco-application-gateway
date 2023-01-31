package handler

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/coco-application-gateway/internal/constants"
	"github.com/mrzack99s/coco-application-gateway/internal/types"
	"github.com/mrzack99s/coco-application-gateway/internal/utils"
)

func return502(ctx *gin.Context) {
	hostname, _ := os.Hostname()
	ctx.JSON(502, gin.H{
		"host":    hostname,
		"message": "Service Unavailable",
	})
}

func Serve(sslEnable bool) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		request := ctx.Request.Clone(context.Background())
		hostname := request.Host
		if strings.Contains(hostname, ":") {
			tmp := strings.Split(hostname, ":")
			hostname = tmp[0]
		}

		path := request.URL.Path
		var endpoint *types.RouteEndpointType

		if sslEnable {
			endpoint = utils.FindEndpointMatchPathMatch(hostname, path, true)
		} else {
			endpoint = utils.FindEndpointMatchPathMatch(hostname, path, false)
		}

		if endpoint != nil {

			switch endpoint.Action {
			case constants.ACTION_PROXY:
				resp, err := utils.HttpJSONRequestWithBytesResponse(
					request.Method,
					fmt.Sprintf("%s%s", endpoint.To, path),
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

				ctx.Data(resp.StatusCode, resp.Header.Get("Content-Type"), buf.Bytes())
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
