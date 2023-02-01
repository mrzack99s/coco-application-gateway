package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/coco-application-gateway/internal/constants"
	"github.com/mrzack99s/coco-application-gateway/internal/utils"
)

func GetGinLog(https bool) gin.HandlerFunc {
	if https {
		return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

			appGwAction := "out-of-service"

			hostname := param.Request.Host
			if strings.Contains(hostname, ":") {
				tmp := strings.Split(hostname, ":")
				hostname = tmp[0]
			}

			msg := ""
			endpoint, foundEndpoint := utils.FindEndpointMatchPathMatch(hostname, param.Path, false)

			if foundEndpoint {

				switch endpoint.Action {
				case constants.ACTION_LOAD_BALANCER:
					appGwAction = constants.ACTION_LOAD_BALANCER
					msg = fmt.Sprintf(", poolName=%s", endpoint.BackendPoolName)
				case constants.ACTION_PROXY:
					appGwAction = constants.ACTION_PROXY
					msg = fmt.Sprintf(", to=%s", endpoint.To)
				case constants.ACTION_REDIRECT:
					appGwAction = constants.ACTION_REDIRECT
					msg = fmt.Sprintf(", to=%s", endpoint.To)
				case constants.ACTION_STRING:
					appGwAction = constants.ACTION_STRING
				}

			}

			return fmt.Sprintf("%-16s timestamp=%v, statusCode=%d, clientIp=%s, method=%s, path=%s >> action=%s%s\n%s",
				"[HTTPS Endpoint]",
				param.TimeStamp.Format(time.RFC3339),
				param.StatusCode,
				param.ClientIP,
				param.Method,
				param.Path,
				appGwAction,
				msg,
				param.ErrorMessage,
			)
		})
	}

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		appGwAction := "out-of-service"

		hostname := param.Request.Host
		if strings.Contains(hostname, ":") {
			tmp := strings.Split(hostname, ":")
			hostname = tmp[0]
		}

		msg := ""
		endpoint, foundEndpoint := utils.FindEndpointMatchPathMatch(hostname, param.Path, false)

		if foundEndpoint {

			switch endpoint.Action {
			case constants.ACTION_LOAD_BALANCER:
				appGwAction = constants.ACTION_LOAD_BALANCER
				msg = fmt.Sprintf(", poolName=%s", endpoint.BackendPoolName)
			case constants.ACTION_PROXY:
				appGwAction = constants.ACTION_PROXY
				msg = fmt.Sprintf(", to=%s", endpoint.To)
			case constants.ACTION_REDIRECT:
				appGwAction = constants.ACTION_REDIRECT
				msg = fmt.Sprintf(", to=%s", endpoint.To)
			case constants.ACTION_STRING:
				appGwAction = constants.ACTION_STRING
			}

		}

		return fmt.Sprintf("%-16s timestamp=%v, statusCode=%d, clientIp=%s, method=%s, path=%s >> action=%s%s\n%s",
			"[HTTP Endpoint]",
			param.TimeStamp.Format(time.RFC3339),
			param.StatusCode,
			param.ClientIP,
			param.Method,
			param.Path,
			appGwAction,
			msg,
			param.ErrorMessage,
		)
	})
}
