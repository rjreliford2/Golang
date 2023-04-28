/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// x001Cmd represents the x001 command
var x001Cmd = &cobra.Command{
	Use:   "x001",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Making API call to cloud endpoint...")

		// Contsructs the API request
		url := "http://localhost:1323/"
		data := []byte(`{"param1":"value1","param2":"value2"}`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			fmt.Println("Error creating API request:", err)
			return
		}

		//Adds any headers to the reqeust
		req.Header.Set("Content-Type", "application/json")

		//Send API request and gets the response
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making API call:", err)
			return
		}
		defer resp.Body.Close()
		//Reads the response and prints it to the console
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading API response:", err)
			return
		}
		fmt.Println("API response:", string(body))
	},
}

func init() {
	rootCmd.AddCommand(x001Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// x001Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// x001Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
