package utils

import (
	"strings"
)

// FormatLocation ensures location strings are correctly formatted
func FormatLocation(location string) string {
	location = strings.ReplaceAll(location, "-", " ")
	location = strings.ReplaceAll(location, "_", " ")
	return location
}
