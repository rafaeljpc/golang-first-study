package util

import "github.com/google/uuid"

func GenerateUUID() string {
	return uuid.New().String()
}

func Substring(s string, start, length int) string {
	if start < 0 || start >= len(s) {
		return ""
	}
	end := start + length
	if end > len(s) {
		end = len(s)
	}
	return s[start:end]
}