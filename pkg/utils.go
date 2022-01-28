package pkg

import (
	"os"
	"sort"
)

func getKeys(data map[string]interface{}) []string {
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func appendKeys(keys []string, add ...string) []string {
	return append(append([]string{}, keys...), add...)
}

func FileIsExist(f string) bool {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false
	}
	return true
}
