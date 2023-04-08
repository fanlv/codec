package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	asciiEncodeCmd = &cobra.Command{
		Use:   "ase [str]",
		Short: "ASCII encode",
		Long:  `ASCII encode`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			asciiEncode(args[0])
		},
	}

	asciiDecodeCmd = &cobra.Command{
		Use:   "asd [str]",
		Short: "ASCII decode",
		Long:  `ASCII decode`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			asciiDecode(args[0])
		},
	}
)

func asciiEncode(str string) {
	fmt.Println("ASCII Encode:")
	fmt.Println(asciiDecodeStr(str))
}

func asciiDecode(str string) {
	fmt.Println("ASCII Decode:")
	fmt.Println(asciiEncodeStr(str))
}

func asciiEncodeStr(str string) string {
	result := ""
	for _, char := range str {
		if int(char) < 128 {
			result += string(char)
		} else {
			result += "\\u" + strconv.FormatInt(int64(char), 16)
		}
	}
	return result
}

func asciiDecodeStr(encodedStr string) string {
	result := ""
	parts := splitString(encodedStr, "\\u")
	for _, part := range parts {
		if len(part) == 4 {
			charCode, err := strconv.ParseInt(part, 16, 32)
			if err != nil {
				fmt.Println("error:", err)
				return ""
			}
			result += string(charCode)
		} else {
			result += part
		}
	}
	return result
}

func splitString(str string, sep string) []string {
	parts := []string{}
	start := 0
	for i := 0; i < len(str)-len(sep)+1; i++ {
		if str[i:i+len(sep)] == sep {
			parts = append(parts, str[start:i])
			start = i + len(sep)
		}
	}
	parts = append(parts, str[start:])
	return parts
}
