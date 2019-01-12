package cmd

import (
	"github.com/alimy/mir"
	"github.com/alimy/mir-music/cmd"
	"github.com/alimy/mir-music/models"
	"github.com/alimy/mir-music/module/serve/openapi"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/unisx/logus"
	"net/http"
	"time"

	ginE "github.com/alimy/mir/module/gin"
)

const (
	listenAddrDefault   = "127.0.0.1:8013" // default listen address
	certFilePathDefault = "cert.pem"       // certificate file default path
	keyFilePathDefault  = "key.pem"        // key file used in https server default path
)

var (
	address     string
	certFile    string
	keyFile     string
	enableHttps bool
	inDebug     bool
)

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "start to mirMusic service",
		Long:  "this cmd will start a https server to provide ginMusic service",
		Run:   serveRun,
	}

	// Parse flags for serveCmd
	serveCmd.Flags().StringVarP(&address, "addr", "a", listenAddrDefault, "service listen address")
	serveCmd.Flags().StringVarP(&certFile, "cert", "c", certFilePathDefault, "certificate path used in https connect")
	serveCmd.Flags().StringVarP(&keyFile, "key", "k", keyFilePathDefault, "key path used in https connect")
	serveCmd.Flags().BoolVarP(&enableHttps, "https", "s", false, "whether use https serve connect")
	serveCmd.Flags().BoolVarP(&inDebug, "debug", "d", false, "whether in debug mode")

	// Register serveCmd as sub-command
	cmd.Register(serveCmd)
}

func serveRun(cmd *cobra.Command, args []string) {
	setup()

	// Instance a default gin engine
	e := gin.Default()

	// Register Api
	registerApi(e)

	// Setup http.Server
	server := &http.Server{
		Handler: e,
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start http.Server
	if enableHttps {
		logus.Info("start listen and serve", logus.String("address", address))
		server.ListenAndServeTLS(certFile, keyFile)
	} else {
		logus.Info("listen and serve",
			logus.String("address", address),
			logus.Bool("enableHttps", enableHttps))
		server.ListenAndServe()
	}
}

func registerApi(e *gin.Engine) {
	// setup mir use *gin.Engine first
	mir.Setup(ginE.Mir(e))
	// register entries to *gin.Engine by mir
	entries := openapi.MirEntries()
	mir.Register(entries...)
}

func setup() {
	if !inDebug {
		logus.InProduction()
		gin.SetMode(gin.ReleaseMode)
	}

	// initial models with MemoryProfile
	models.Register(models.MemoryProfile)
}
