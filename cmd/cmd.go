package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codec",
	Short: "常用的字符串编解码",
}

func Execute() {
	rootCmd.AddCommand(urlEncodeCmd)
	rootCmd.AddCommand(urlDecodeCmd)

	rootCmd.AddCommand(utf8EncodeCmd)
	rootCmd.AddCommand(utf8DecodeCmd)

	rootCmd.AddCommand(unicodeEncodeCmd)
	rootCmd.AddCommand(unicodeDecodeCmd)

	rootCmd.AddCommand(base64EncodeCmd)
	rootCmd.AddCommand(base64DecodeCmd)

	rootCmd.AddCommand(base64IntEncodeCmd)
	rootCmd.AddCommand(base64IntDecodeCmd)

	rootCmd.AddCommand(asciiEncodeCmd)
	rootCmd.AddCommand(asciiDecodeCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
