package cfg_reader

import "strings"

func Unescape(str string) string {
	res := ""
	end := len(str) - 1
	for pos, char := range str {
		c := string(char)
		if c == "\\" && pos != end {
			c = string(str[pos+1])
			switch c {
			case "n":
			case "\n":
				res += "\n"
			case "r":
				res += "\r"
			case "'":
			case "\"":
			case "\\":
				res += c
			default:
				res += "\\"
				res += c
			}

			continue
		}

		res += c
	}

	res = strings.TrimSpace(res)

	return res
}

func Escape(str string) string {
	res := ""

	for _, char := range str {
		c := string(char)

		switch c {
		case "\n":
			res += "\\n"
		case "\r":
			res += "\\r"
		case "'":
		case "\"":
		case "\\":
			res += "\\"
			res += c
		default:
			res += c
		}
	}

	return res
}
