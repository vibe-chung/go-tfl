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

// crowdingCmd represents the crowding command
var crowdingCmd = &cobra.Command{
	Use:   "crowding",
	Short: "Get crowding information for a specific stop point (naptan) from the TFL API.",
	Long: `Fetches and displays crowding information for a given London Underground stop point (naptan) using the Transport for London (TFL) public API.

Usage:
  go-tfl crowding <naptan> [--live]

Example:
  go-tfl crowding 940GZZLUBND
  go-tfl crowding 940GZZLUBND --live

The --live flag will fetch live crowding data if available. The command returns formatted JSON data for the specified stop point.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		naptan := args[0]
		live, _ := cmd.Flags().GetBool("live")
		var url string
		if live {
			url = fmt.Sprintf("https://api.tfl.gov.uk/Crowding/%s/live", naptan)
		} else {
			url = fmt.Sprintf("https://api.tfl.gov.uk/Crowding/%s", naptan)
		}
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
		var prettyObj map[string]interface{}
		if err := json.Unmarshal(body, &prettyObj); err == nil {
			out, _ := json.MarshalIndent(prettyObj, "", "  ")
			fmt.Println(string(out))
			return
		}
		fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(crowdingCmd)
	crowdingCmd.Flags().Bool("live", false, "Fetch live crowding data if available")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crowdingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crowdingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
