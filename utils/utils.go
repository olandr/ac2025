package utils

import "strconv"

func Int64(v string) int64 {
	if r, err := strconv.ParseInt(v, 0, 64); err == nil {
		return r
	}
	return -1
}
