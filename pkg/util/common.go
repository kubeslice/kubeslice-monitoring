package util

import "strings"

func MergeMaps(ms ...map[string]string) map[string]string {
	res := map[string]string{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

func FlattenKey(key string) string {
	key = strings.Replace(key, " ", "_", -1)
	key = strings.Replace(key, ".", "_", -1)
	key = strings.Replace(key, "-", "_", -1)
	key = strings.Replace(key, "=", "_", -1)
	key = strings.Replace(key, "/", "_", -1)
	return key
}

func FlattenMap(m map[string]string) map[string]string {
	res := make(map[string]string)
	for k, v := range m {
		res[FlattenKey(k)] = FlattenKey(v)
	}
	return res
}
