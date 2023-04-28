/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"server/handlers"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// modifyjsonCmd represents the modifyjson command
var modifyjsonCmd = &cobra.Command{
	Use:   "modifyjson",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		baseurl, _ := cmd.Flags().GetString("baseurl")
		port, _ := cmd.Flags().GetString("port")
		endpoint, _ := cmd.Flags().GetString("endpoint")
		file, _ := cmd.Flags().GetString("file")
		url := baseurl + ":" + port + endpoint
		responsedata, err := handlers.SendJson(url, file)
		if err != nil {
			fmt.Println(err)
			return
		}
		var cleanJson bytes.Buffer
		err = json.Indent(&cleanJson, responsedata, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.WriteFile("response.json", cleanJson.Bytes(), 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("modifyjson called")
	},
}

func init() {
	rootCmd.AddCommand(modifyjsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modifyjsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modifyjsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	x002Cmd.Flags().StringP("endpoint", "l", "end", "URL endpoint")
	x002Cmd.Flags().StringP("port", "p", "1323", "HTTP port number")
	x002Cmd.Flags().StringP("baseurl", "u", "localhost", "HTTP url")
	x002Cmd.Flags().StringP("file", "f", "filepath", "file path of the json")
	viper.BindPFlag("baseurl", x002Cmd.Flags().Lookup("baseurl"))
	viper.BindPFlag("port", x002Cmd.Flags().Lookup("port"))
	viper.BindPFlag("endpoint", x002Cmd.Flags().Lookup("endpoint"))
	viper.BindPFlag("file", x002Cmd.Flags().Lookup("file"))
}
