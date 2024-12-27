/*
Copyright © 2024 Steven Davis
*/
package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Logs all requests",
	Long: `	Logs all requests as either SENT or RECEIVED 
	

	Provides brief description of request.
	
		`,

	Run: func(cmd *cobra.Command, args []string) {

		faustPath, _ := os.UserHomeDir()

		faustPath = fmt.Sprintf("%s/faust/log/", faustPath)

		output, _ := cmd.Flags().GetString("output")
		remove, _ := cmd.Flags().GetString("remove")

		if output != "" {
			faustPath = fmt.Sprintf("%s%s", faustPath, output)
			fmt.Println(faustPath)
			if _, err := os.Stat(faustPath); os.IsNotExist(err) {
				fmt.Println("Log does not exist - check file name")
				return
			} else {
				data, _ := os.ReadFile(faustPath)
				fmt.Print(string(data))
				return
			}
		}

		if remove != "" {
			faustPath = fmt.Sprintf("%s%s", faustPath, remove)
			if _, err := os.Stat(faustPath); os.IsNotExist(err) {
				fmt.Println("Log does not exist - check file name")
				return
			} else {
				err := os.Remove(faustPath)
				if err != nil {
					panic(err)
				}
				return
			}
		}

		files, err := os.ReadDir(faustPath)

		if err != nil {
			fmt.Println("Error reading directory: ", err)
			return
		}
		sort.Slice(files, func(i, j int) bool {
			return files[i].Name() > files[j].Name()
		})

		for _, file := range files {
			fmt.Println(file.Name())
		}

	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	logCmd.Flags().StringP("output", "o", "", "Outputs the log into the terminal")
	logCmd.Flags().StringP("remove", "r", "", "Deletes the log with matching file name")

}
