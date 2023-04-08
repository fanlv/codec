package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var (
	base64EncodeCmd = &cobra.Command{
		Use:   "base64e [str]",
		Short: "Base64 encode",
		Long:  `Base64 encode`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			base64Encode(args[0])
		},
	}

	base64DecodeCmd = &cobra.Command{
		Use:   "base64d [str]",
		Short: "Base64 decode",
		Long:  `Base64 decode`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			base64Decode(args[0])
		},
	}

	base64IntEncodeCmd = &cobra.Command{
		Use:   "base64ie [str]",
		Short: "Base64 int encode",
		Long:  `Base64 int encode`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			base64IntEncode(args[0])
		},
	}

	base64IntDecodeCmd = &cobra.Command{
		Use:   "base64id [str]",
		Short: "Base64 int decode",
		Long:  `Base64 int decode`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			base64IntDecode(args[0])
		},
	}
)

func base64Encode(str string) {
	fmt.Println("Base64 Encode:")
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(str)))
}

func base64IntEncode(str string) {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("ParseInt Error:", err.Error())
		return
	}
	fmt.Println("Base64 Int Encode:")
	fmt.Println(encodeLogKey(num))
}

func base64IntDecode(str string) {
	res, err := decodeLogKey(str)
	if err != nil {
		fmt.Println("Base64 Int Decode Error:", err.Error())
		return
	}
	fmt.Println("Base64 Int Decode:")
	fmt.Println(res)
}

func base64Decode(str string) {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("Base64 Decode error:", err)
		return
	}

	fmt.Println("Base64 Decode:")
	fmt.Println(string(decoded))
}

var codes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"

func encodeLogKey(id int64) string {
	str := make([]byte, 0, 12)
	if id == 0 {
		return "0"
	}
	for id > 0 {
		ch := codes[id%64]
		str = append(str, byte(ch))
		id /= 64
	}
	return string(str)
}

func decodeLogKey(logKey string) (int64, error) {
	res := int64(0)

	for i := len(logKey); i > 0; i-- {
		ch := logKey[i-1]
		res *= 64
		mod := strings.IndexRune(codes, rune(ch))
		if mod == -1 {
			return -1, fmt.Errorf("Invalid logKey character: '%c'", ch)
		}
		res += int64(mod)
	}
	return res, nil
}
