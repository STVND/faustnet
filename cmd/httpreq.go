/*
Copyright © 2024 Steven Davis
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var export bool

var reqCmd = &cobra.Command{
	Use:   "httpreq",
	Short: "Create HTTP Requests",
	Long:  `Allows user to create HTTP Requests.`,

	Run: func(cmd *cobra.Command, args []string) {
		faustPath := viper.GetString("FaustDir")
		hosting := viper.GetString("HostedSubDir")

		uri, _ := cmd.Flags().GetString("uri")
		method, _ := cmd.Flags().GetString("method")
		bearer, _ := cmd.Flags().GetString("bearer")
		export, _ := cmd.Flags().GetBool("export")
		filepath, _ := cmd.Flags().GetString("filepath")

		startTime := time.Now().Format("2006-01-02_15-04-05")

		switch strings.ToLower(method) {
		case "delete":
			resp, err := http.NewRequest("DELETE", uri, nil)

			if err != nil {
				fmt.Printf("Error creating HTTP request: %s\n", err)
				os.Exit(1)
			}

			httpResp := FormatSendRequest(resp, bearer)
			fmt.Println(httpResp.Status)

			defer httpResp.Body.Close()

		case "post":
			if filepath == "" {
				fmt.Println("No file selected - recommend using --filepath flag to set filepath")
				os.Exit(1)
			}

			var filecontents []byte
			_ = filecontents

			if _, err := os.Stat(filepath); err != nil {
				tryPath := faustPath + "/" + hosting + "/" + filepath
				if _, err := os.Stat(tryPath); err != nil {
					fmt.Println("Unable to locate path to complete request")
					os.Exit(1)
				}
				filecontents, _ = os.ReadFile(tryPath)
			} else {
				filecontents, _ = os.ReadFile(filepath)
			}

			resp, err := http.NewRequest("POST", uri, bytes.NewReader(filecontents))
			if err != nil {
				fmt.Println("Error creating request")
				os.Exit(1)
			}

			httpResp := FormatSendRequest(resp, bearer)
			extension := GetExtension(filepath)
			contentHeader := mime.TypeByExtension(extension)
			resp.Header.Add("Content-Type", contentHeader)
			print(httpResp.Status)

			defer resp.Body.Close()

		case "get":
			resp, err := http.NewRequest("GET", uri, nil)

			if err != nil {
				fmt.Printf("Error creating HTTP request: %s\n", err)
				os.Exit(1)
			}

			httpResp := FormatSendRequest(resp, bearer)

			defer httpResp.Body.Close()

			contentType := httpResp.Header.Get("Content-Type")

			body, err := io.ReadAll(httpResp.Body)
			if err != nil {
				fmt.Println("Error reading response body: ", err)
				os.Exit(1)
			}

			extension, err := mime.ExtensionsByType(contentType)
			if err != nil {
				fmt.Println("Error returning file extension: ", err)
				extension = []string{".txt"}
			}

			fmt.Println(string(body))

			if export {

				Export(faustPath, startTime, body, extension[0])
			}

		default:
			fmt.Println("Unable to understand HTTP Request type - possibly not available")
			os.Exit(1)
		}

	},
}

func FormatSendRequest(resp *http.Request, bearer string) *http.Response {

	if bearer != "" {
		resp.Header.Add("Authorization", "Bearer "+bearer)
	}

	resp.Header.Add("Accept", "application/json, application/yaml, application/xml, text/csv, text/html")

	client := http.Client{}
	httpResp, err := client.Do(resp)

	if err != nil {
		fmt.Printf("Error sending HTTP request: %s\n", err)
		os.Exit(1)
	}

	return httpResp
}

func GetExtension(filepath string) string {
	values := strings.Split(filepath, ".")
	return "." + values[len(values)-1]
}

func Export(faustPath string, startTime string, body []byte, ext string) {
	exportPath := fmt.Sprintf("%s/received/%s%s", faustPath, startTime, ext)
	fmt.Println(exportPath)
	_, err := os.Create(exportPath)
	if err != nil {
		fmt.Printf("Error creating file: %s", err)
	}
	err = os.WriteFile(exportPath, body, 0644)
	if err != nil {
		fmt.Printf("Error writing to %s file: \n", ext)
		fmt.Sprintln(err)
		os.Exit(1)
	}

}

func init() {
	rootCmd.AddCommand(reqCmd)
	// Here you will define your flags and configuration settings.
	reqCmd.Flags().StringP("uri", "u", "", "Destination uri")
	reqCmd.Flags().StringP("method", "m", "get", "Desired HTTP Method")
	reqCmd.Flags().StringP("bearer", "b", "", "Bearer token - if expected")
	reqCmd.Flags().BoolVarP(&export, "export", "e", false, "Saves to local drive")
	reqCmd.MarkFlagRequired("uri")
	reqCmd.Flags().StringP("filepath", "f", "", "Used with put command to select file to be uploaded. Attempts as absolute, then attempts in hosting folder")
}
