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
	Long:  `Query the information of your ip or the specific ip.
For example:
	devbox ip           query the ip info of your host.

	devbox ip 8.8.8.8   query the specific ip info.
`,
	Run: func(cmd *cobra.Command, args []string) {
		var ip string
		if len(args) == 1 {
			ip = args[0]
		} else if len(args) > 1 {
			fmt.Println("error args")
		}
		if err := printIPv4Info(ip); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}

func printIPv4Info(ip string) error {
	var url string
	if ip == "" {
		url = "https://ipinfo.io"
	} else {
		url = "https://ipinfo.io" + "/" + ip
	}
	rsp, err := http.Get(url)
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
