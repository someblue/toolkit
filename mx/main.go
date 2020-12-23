package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use:  "mx",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addr := args[0]
		mxs, err := net.LookupMX(addr)
		if err != nil {
			fmt.Printf("failed to lookup mx [%v]: %v\n", addr, err)
			os.Exit(1)
		}
		for _, mx := range mxs {
			fmt.Printf("pref: %v, host: %v\n", mx.Pref, mx.Host)
			ips, err := net.LookupIP(mx.Host)
			if err != nil {
				fmt.Printf("failed to lookup ip [%v]: %v\n", mx.Host, err)
				continue
			}
			for _, ip := range ips {
				if ipv4 := ip.To4(); ipv4 == nil {
					fmt.Printf("\t%v (ipv6)\n", ip)
					continue
				}

				resp, err := http.Get(fmt.Sprintf("https://api.ip138.com/ipv4/?ip=%s&datatype=txt&token=%s", ip, os.Getenv("IP138TOKEN")))
				if err != nil {
					fmt.Printf("\tfailed to query ip138 for ip [%v]: %v", ip, err)
					continue
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("\tfailed to read resp body after query ip138 for ip [%v]: %v", ip, err)
					continue
				}
				fmt.Printf("\t%s\n", body)
			}
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
