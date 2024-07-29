package helpers

import "strings"

func CreateUrl(title string) string {
	title = strings.ToLower(title)
	url := ""
	for _, char := range title {
		switch char {
		case 'ö':
			url += "o"
		case 'ü':
			url += "u"
		case 'ğ':
			url += "g"
		case 'ç':
			url += "c"
		case 'ı':
			url += "i"
		case 'ş':
			url += "s"
		case ' ':
			url += "-"
		case '?':
			url += "-"

		default:
			url += string(char)
		}
	}
	return url
}
