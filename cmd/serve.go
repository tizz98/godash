package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/tizz98/godash/api"
	"github.com/tizz98/godash/app"
	"log"
	"net/http"
)

var port int

func init() {
	serve.PersistentFlags().IntVarP(&port, "port", "p", 9090, "the port to run the server on")

	rootCmd.AddCommand(serve)
}

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Starts the http web server",
	Run: func(cmd *cobra.Command, args []string) {
		app := app.NewApp()
		api := api.NewApi(app)

		r := mux.NewRouter()
		r.HandleFunc("/stocks", api.SearchStocks).Methods("GET")

		app.Context.Logger.Info(fmt.Sprintf("Started server at http://localhost:%d", port))
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
	},
}
