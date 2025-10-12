/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// stopsCmd represents the stops command
var stopsCmd = &cobra.Command{
	Use:   "stops",
	Short: "Get all stop points for a specific London Underground line from the TFL API.",
	Long: `Fetches and displays all stop points (stations) for a given London Underground line using the Transport for London (TFL) public API.

Usage:
  go-tfl line stops <line>

Example:
  go-tfl line stops central

This command will return a formatted JSON array of stop points for the specified line.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lineArg := args[0]
		url := fmt.Sprintf("https://api.tfl.gov.uk/Line/%s/StopPoints", lineArg)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			return
		}
		req.Header.Set("User-Agent", "curl/8.7.1")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error fetching data: %v\n", err)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			return
		}
		var prettyArr []interface{}
		if err := json.Unmarshal(body, &prettyArr); err == nil {
			out, _ := json.MarshalIndent(prettyArr, "", "  ")
			fmt.Println(string(out))
			return
		}
		fmt.Println(string(body))
	},
}

func init() {
	lineCmd.AddCommand(stopsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
