package core

import (
	"fmt"
	"regexp"
	"strings"
)

func Hello() {
	fmt.Println("Hello, World!")
}

func FilterSlice(slice []string, filters ...func(string) bool) []string {
	var result []string
	for _, item := range slice {
		include := true
		for _, filter := range filters {
			if !filter(item) {
				include = false
				break
			}
		}
		if include {
			result = append(result, item)
		}
	}
	return result
}

func IncludeContains(substrings []string) func(string) bool {
	return func(item string) bool {
		for _, substring := range substrings {
			if strings.Contains(strings.ToLower(item), strings.ToLower(substring)) {
				return true
			}
		}
		return false
	}
}

func ExcludeContains(substrings []string) func(string) bool {
	return func(item string) bool {
		for _, substring := range substrings {
			if strings.Contains(strings.ToLower(item), strings.ToLower(substring)) {
				return false
			}
		}
		return true
	}
}

func IncludeMatches(patterns []string) func(string) bool {
	return func(item string) bool {
		for _, pattern := range patterns {
			if matched, _ := regexp.MatchString("(?i)"+pattern, item); matched {
				return true
			}
		}
		return false
	}
}

func ExcludeMatches(patterns []string) func(string) bool {
	return func(item string) bool {
		for _, pattern := range patterns {
			if matched, _ := regexp.MatchString("(?i)"+pattern, item); matched {
				return false
			}
		}
		return true
	}
}

