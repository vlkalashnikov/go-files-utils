package files

import "strings"

func GetFileExtension(fileName string) string {
	arr := strings.Split(fileName, ".")
	if len(arr) == 1 {
		return ""
	}
	return arr[len(arr)-1]
}
