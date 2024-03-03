package support

import (
	"fmt"
	"strings"
)

func Contains[T comparable](array []T, element T) bool {
	for _, e := range array {
		if element == e {
			return true
		}
	}
	return false
}

func FormatString(str string, params []map[string]interface{}) string {
	for _, values := range params {
		for key, value := range values {
			placeholder := "{" + key + "}"
			str = replacePlaceholder(str, placeholder, value)
		}
	}
	return str
}

func replacePlaceholder(str, placeholder string, value interface{}) string {
	return strings.ReplaceAll(str, placeholder, fmt.Sprintf("%v", value))
}
