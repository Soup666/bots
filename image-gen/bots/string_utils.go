package bots

type StringUtils struct{}

func NewStringUtils() *StringUtils {
	return &StringUtils{}
}

// StringUtils provides utility functions for string manipulation
func (su *StringUtils) StringUtils(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}
