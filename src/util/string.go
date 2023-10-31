package util

func StrDefault(strToCheck string, defaultStr string) string {
	if strToCheck == "" {
		return defaultStr
	}
	return strToCheck
}
