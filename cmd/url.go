package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/url"
)

var (
	urlEncodeCmd = &cobra.Command{
		Use:   "ue [str]",
		Short: "url encode",
		Long:  `url encode`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			urlEncode(args[0])
		},
	}

	urlDecodeCmd = &cobra.Command{
		Use:   "ud [str]",
		Short: "url decode",
		Long:  `url decode`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			urlDecode(args[0])
		},
	}
)

func urlEncode(str string) {
	fmt.Println("URL Encode:")
	fmt.Println(url.QueryEscape(str))
}

func urlDecode(str string) {
	res, err := url.QueryUnescape(str)
	if err != nil {
		fmt.Println("decode err = ", err)
		return
	}

	fmt.Println("URL Decode:")
	fmt.Println(res)
}
