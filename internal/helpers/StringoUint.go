package helpers

import (
    "strconv"
    "errors"
)

// StringToUint converts a string to uint, returns error if conversion fails.
func StringToUint(s string) (uint, error) {
    i, err := strconv.ParseUint(s, 10, 64)
    if err != nil {
        return 0, errors.New("invalid uint string")
    }
    return uint(i), nil
}
