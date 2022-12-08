package util

import (
	"strconv"
)

// ParseUint .
func ParseUint(key string) (uint, error) {
	if key == "" {
		return 0, nil
	}
	id, err := strconv.ParseUint(key, 10, strconv.IntSize)
	return uint(id), err
}
