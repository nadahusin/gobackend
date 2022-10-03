package config

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nadahusin/gorent/src/routers"

	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "start",
	Short: "Application Start",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {

		var addrs string = "0.0.0.0:8082"

		if pr := os.Getenv("PORT"); pr != "" {
			addrs = "0.0.0.0:" + pr
		}
		srv := &http.Server{
			Addr:         addrs,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      mainRoute,
		}
		fmt.Println("Application Run on http://", addrs)
		srv.ListenAndServe()
		return nil
	} else {
		return err
	}
}
