package internal

import "strings"

// Unescape string
func Unescape(str string) string {
	res := ""
	end := len(str) - 1

	for p, c := range str {
		if c == '\\' && p != end {
			c = rune(str[p+1])
			switch c {
			case 'n', '\n':
				res += "\n"
			case 'r':
				res += "\r"
			case '\'', '"', '\\':
				res += string(c)
			default:
				res += "\\" + string(c)
			}

			continue
		}

		res += string(c)
	}

	res = strings.TrimSpace(res)

	return res
}

// Escape string
func Escape(str string) string {
	res := ""

	for _, c := range str {

		switch c {
		case '\n':
			res += "\\n"
		case '\r':
			res += "\\r"
		case '\\', '"', '\'':
			res += "\\"
			res += string(c)
		default:
			res += string(c)
		}
	}

	return res
}
