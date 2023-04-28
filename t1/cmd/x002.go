/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"t1/cmd/logic"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// x002Cmd represents the x002 command
var x002Cmd = &cobra.Command{
	Use:   "x002",
	Short: "This is command x002",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("x002 called")
		logic.Z01()
		url, _ := cmd.Flags().GetString("url")
		port, _ := cmd.Flags().GetString("port")
		res, err := z02(url, port)
		fmt.Println(err)
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(x002Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// x002Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// x002Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//x002Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//Add URL, Port, and Endpoint flags
	x002Cmd.Flags().StringP("endpoint", "l", "end", "URL endpoint")
	x002Cmd.Flags().StringP("port", "p", "1323", "HTTP port number")
	x002Cmd.Flags().StringP("url", "u", "localhost", "HTTP url")
	//Bind the flags
	viper.BindPFlag("url", x002Cmd.Flags().Lookup("url"))
	viper.BindPFlag("port", x002Cmd.Flags().Lookup("port"))
	viper.BindPFlag("endpoint", x002Cmd.Flags().Lookup("endpoint"))
}

/*func z01() {
	fmt.Println("z01")
	url := "https:cloundenpointgoeshere.com"
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading API response:", err)
		return
	}
	fmt.Println("API response:", string(body))
}*/

func z02(url string, port string) (string, string) {
	res := fmt.Sprintf(url + port)
	//res := "yada"
	return res, "Hi there"
}
