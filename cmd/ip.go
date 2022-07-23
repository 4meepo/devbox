package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Query the information of your ip or the specific ip. ",
	Long:  `Query the information of your ip or the specific ip.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := queryMyPublicIPv4Info()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}

func queryMyPublicIPv4Info() error {
	rsp, err := http.Get("https://ipinfo.io/json")
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}
