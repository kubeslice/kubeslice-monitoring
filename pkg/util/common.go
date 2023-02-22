package util

import "strings"

// MergeMaps merges multiple maps of string into one
func MergeMaps(ms ...map[string]string) map[string]string {
	res := map[string]string{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// KeysFromMap returns the keys of a map of string as a slice of string
func KeysFromMap(ms map[string]string) []string {
	var res []string
	for k := range ms {
		res = append(res, k)
	}
	return res
}

// FlattenString replaces spaces, dots, dashes, equal signs and slashes with underscores in a string
func FlattenString(key string) string {
	key = strings.Replace(key, " ", "_", -1)
	key = strings.Replace(key, ".", "_", -1)
	key = strings.Replace(key, "-", "_", -1)
	key = strings.Replace(key, "=", "_", -1)
	key = strings.Replace(key, "/", "_", -1)
	return key
}

// FlattenMap replaces spaces, dots, dashes, equal signs and slashes with underscores in a map of string
func FlattenMap(m map[string]string) map[string]string {
	res := make(map[string]string)
	for k, v := range m {
		res[FlattenString(k)] = FlattenString(v)
	}
	return res
}
