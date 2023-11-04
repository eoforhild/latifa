package cmd

import (
	"context"
	"errors"
	"fmt"
	"latifa/config"
	"latifa/router"
	log2 "log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/NYTimes/logrotate"
	"github.com/apex/log"
	"github.com/apex/log/handlers/multi"
	"github.com/apex/log/handlers/text"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "carbon",
	PreRun: func(cmd *cobra.Command, args []string) {
		initConfig()
		initLogging()
		//initDb()
	},
	Run: rootCmdRun,
}

var (
	debug       = false
	configPath  = config.DefaultLocation
	useAutoTls  = false
	tlsHostname = ""
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "run in debug mode")
	rootCmd.PersistentFlags().StringVar(&configPath, "config", config.DefaultLocation, "set the location for the config file")
	rootCmd.PersistentFlags().BoolVar(&useAutoTls, "auto-tls", false, "generate and manage own SSL certificates using Let's Encrypt")
	rootCmd.PersistentFlags().StringVar(&tlsHostname, "tls-hostname", "", "the FQDN for the generated SSL certificate")

	//rootCmd.AddCommand(versionCmd)
}

func rootCmdRun(cmd *cobra.Command, _ []string) {
	r := router.NewClient()

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%d", config.Get().Api.Host, config.Get().Api.Port),
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.WithField("error", err).Fatal("an error occured while trying to start the webserver")
		}
	}()

	log.Info("webserver started, starting context")

	q := make(chan os.Signal, 1)
	// Wait and accept graceful shutdowns when quit via SIGINT (Ctrl+C or DEL)
	// SIGKILL, SIGQUIT, or SIGTERM will not be caught.
	signal.Notify(q, os.Interrupt)

	// Block until a signal is received.
	<-q

	log.Warn("shutting down, waiting for context")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// Catch any errors when closing listeners.
	if err := s.Shutdown(ctx); err != nil {
		panic(err)
	}

	// Since we don't have to wait for any other services to finalize, we don't
	// need to block on <-ctx.Done(). It may be needed in the future.
	os.Exit(0)
}

func initConfig() {
	err := config.FromFile(configPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			exitWithConfigurationError()
		}
		log2.Fatal("cmd/root: failed to create the configuration file: ", err)
	}
}

func initLogging() {
	d := config.Get().LogDirectory
	log.SetLevel(log.InfoLevel)
	if debug || config.Get().Debug {
		log.SetLevel(log.DebugLevel)
	}

	p := filepath.Join(d, "/carbon.log")
	w, err := logrotate.NewFile(p)
	if err != nil {
		log2.Fatal("cmd/root: failed to create log file: ", err)
	}

	log.SetHandler(multi.New(
		text.New(os.Stderr),
		text.New(w),
	))
}

func exitWithConfigurationError() {
	log2.Fatal("The configuration file could not be read or found. Try running with --config.")
	os.Exit(1)
}
