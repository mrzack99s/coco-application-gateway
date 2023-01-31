package main

import (
	"github.com/gin-gonic/gin"
	appwgw_cmd "github.com/mrzack99s/coco-application-gateway/cmd"
	"github.com/mrzack99s/coco-application-gateway/internal/utils"
	"github.com/spf13/cobra"
)

func main() {

	var cmdRun = &cobra.Command{
		Use:   "run",
		Short: "To run an application",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			gin.SetMode(gin.ReleaseMode)
			appwgw_cmd.RunAppGW()
		},
	}

	var cmdCertificate = &cobra.Command{
		Use:   "gencert",
		Short: "To generate a self-signed certificate.",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			utils.GenerateSelfSignCert()
		},
	}

	var rootCmd = &cobra.Command{Use: "coco"}
	rootCmd.AddCommand(cmdRun)
	rootCmd.AddCommand(cmdCertificate)
	rootCmd.Execute()
}
