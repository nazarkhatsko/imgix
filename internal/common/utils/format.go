package utils

import "fmt"

func GetFormatPageTitle(appName, pageName string) string {
	return fmt.Sprintf("%s | %s", appName, pageName)
}
