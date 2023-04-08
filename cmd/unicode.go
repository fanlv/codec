package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
	"strconv"
	"unicode/utf8"
)

var (
	unicodeEncodeCmd = &cobra.Command{
		Use:   "unie [str]",
		Short: "Unicode encode",
		Long:  `Unicode encode`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			unicodeEncode(args[0])
		},
	}

	unicodeDecodeCmd = &cobra.Command{
		Use:   "unid [str]",
		Short: "Unicode decode",
		Long:  `Unicode decode`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			unicodeDecode(args[0])
		},
	}
)

func unicodeEncode(str string) {
	fmt.Println("Unicode Encode:")
	fmt.Println(unicodeEncodeStr(str))
}

func unicodeDecode(str string) {
	fmt.Println("URL Decode:")
	fmt.Println(unicodeDecodeStr(str))
}

func unicodeEncodeStr(str string) string {
	var encoded []byte

	for _, r := range str {
		encoded = append(encoded, []byte(fmt.Sprintf("&#%d;", r))...)
	}

	return string(encoded)
}

func unicodeDecodeStr(encodedStr string) string {
	re := regexp.MustCompile(`&#(\d+);`)
	decoded := re.ReplaceAllStringFunc(encodedStr, func(match string) string {
		dec, _ := strconv.Atoi(match[2 : len(match)-1])
		return string(rune(dec))
	})
	decodedBytes := []byte(decoded)

	if !utf8.Valid(decodedBytes) {
		decodedBytes = []byte(utf8ReplaceInvalid(decodedBytes))
	}

	return string(decodedBytes)
}
