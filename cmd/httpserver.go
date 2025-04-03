/*
Copyright © 2024 Steven Davis
*/
package cmd

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "httpserv",
	Short: "Manage a simple website to host your files",

	Run: func(cmd *cobra.Command, args []string) {

		bearer, _ := cmd.Flags().GetString("bearer")
		port, _ := cmd.Flags().GetString("port")

		if port != "" {

		}

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("Unable to read configuration file: %s\n", err)
			os.Exit(1)
		}

		startTime := time.Now()
		serverpath := (viper.GetString("FaustDir") + "/" + viper.GetString("HostedSubDir"))
		fmt.Println("Attempting to start server")

		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			allowed := checkBearer(r, bearer)

			if allowed {

				w.Header().Add("Content-Type", "text/plain")
				dirContents, err := os.ReadDir(serverpath)
				if err != nil {
					fmt.Fprintf(w, "Sorry! Unable to retrieve contents: %s", err)
				}
				fmt.Fprintf(w, "Available Files:\n\n")
				for _, entry := range dirContents {
					fmt.Fprintf(w, "%s\n", entry.Name())
				}
			} else {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}
		})

		cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
		if err != nil {
			fmt.Println("Error creating tls: ", err)
		}
		tlsconf := &tls.Config{Certificates: []tls.Certificate{cer}}

		server := &http.Server{
			Addr:      ":8080",
			Handler:   mux,
			TLSConfig: tlsconf,
		}

		go func() {
			fmt.Println("Starting server")

			if err := server.ListenAndServe(); err != http.ErrServerClosed {
				fmt.Println("Error starting server:", err)
			}
		}()

		go func() {
			for {
				uptime := time.Since(startTime)

				log.Printf("Server uptime: %.4s", uptime)
				time.Sleep(10 * time.Second)

			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		fmt.Println("Shutting down server")

		if err := server.Close(); err != nil {
			fmt.Println("Error shutting down server:", err)
		}

		fmt.Println("Server stopped")

	},
}

func checkBearer(r *http.Request, br string) bool {
	//naive way to do a bearer check
	//but for the scope of this project it should be secure *enough*

	getReqBearer := r.Header.Get("Authorization")

	return getReqBearer == br

}

func init() {
	rootCmd.AddCommand(serverCmd)
	//naive assignment of bearer token but this is a small project
	serverCmd.Flags().StringP("bearer", "b", "", "Set a fixed bearer token for user access")
	serverCmd.Flags().StringP("port", "p", "8080", "Specify port number to use")
}
