package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
	"strconv"
	"unicode/utf8"
)

var (
	utf8EncodeCmd = &cobra.Command{
		Use:   "u8e [str]",
		Short: "utf8 encode",
		Long:  `utf8 encode`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			utf8Encode(args[0])
		},
	}

	utf8DecodeCmd = &cobra.Command{
		Use:   "u8d [str]",
		Short: "utf8 decode",
		Long:  `utf8 decode`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			utf8Decode(args[0])
		},
	}
)

func utf8Encode(str string) {
	var encoded []byte

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]

		if r >= 128 {
			encoded = append(encoded, []byte(fmt.Sprintf("&#x%X;", r))...)
		} else {
			encoded = append(encoded, byte(r))
		}
	}

	fmt.Println("UTF8 Encode :")
	fmt.Println(string(encoded))
}

func utf8Decode(str string) {
	fmt.Println("UTF8 Encode :")
	fmt.Println(utf8DecodeStr(str))
}

func utf8DecodeStr(encodedStr string) string {
	re := regexp.MustCompile(`&#x([0-9A-Fa-f]+);`)
	decoded := re.ReplaceAllStringFunc(encodedStr, func(match string) string {
		hex := match[3 : len(match)-1]
		dec, _ := strconv.ParseInt(hex, 16, 32)
		return strconv.FormatInt(dec, 10)
	})

	decodedBytes := []byte(decoded)

	if !utf8.Valid(decodedBytes) {
		decodedBytes = []byte(utf8ReplaceInvalid(decodedBytes))
	}

	return string(decodedBytes)
}

func utf8ReplaceInvalid(input []byte) string {
	output := make([]rune, 0, len(input))
	for len(input) > 0 {
		r, size := utf8.DecodeRune(input)
		input = input[size:]
		if r == utf8.RuneError {
			output = append(output, '�')
		} else {
			output = append(output, r)
		}
	}
	return string(output)
}
