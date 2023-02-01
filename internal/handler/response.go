package handler

import (
	"encoding/json"
	"os"
	"time"

	"github.com/corazawaf/coraza/v2/loggers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mrzack99s/coco-application-gateway/internal/logger"
)

func return502(ctx *gin.Context) {
	hostname, _ := os.Hostname()
	ctx.JSON(502, gin.H{
		"host":    hostname,
		"message": "Service Unavailable",
	})
}

func return500(ctx *gin.Context, msg string) {
	hostname, _ := os.Hostname()
	ctx.JSON(500, gin.H{
		"host":    hostname,
		"message": msg,
	})
}

func return403(ctx *gin.Context, msg string) {
	hostname, _ := os.Hostname()
	ctx.JSON(502, gin.H{
		"host":    hostname,
		"message": msg,
	})
}

func returnWAFforbidden(ctx *gin.Context, path string, auditLog *loggers.AuditLog, https bool) {
	hostname, _ := os.Hostname()
	headName := "[HTTP Endpoint]"
	if https {
		headName = "[HTTPS Endpoint]"
	}
	msg, _ := json.MarshalIndent(auditLog.Messages, "", "\t")

	uuid, _ := uuid.NewUUID()

	logger.WAFLogger.Printf("%-16s timestamp=%v, clientIp=%s, method=%s, path=%s, action=%s\n>> Start Track ID %s\n%v\n>> End of trackId %s \n",
		headName,
		time.Now().Format(time.RFC3339),
		ctx.ClientIP(),
		ctx.Request.Method,
		path,
		"deny",
		uuid.String(),
		string(msg),
		uuid.String(),
	)

	ctx.JSON(403, gin.H{
		"host":    hostname,
		"status":  "Sorry, you have been blocked",
		"action":  "deny",
		"trackId": uuid.String(),
	})

}

func response(ctx *gin.Context, statusCode int, contentType string, data []byte) {
	ctx.Data(statusCode, contentType, data)
}
