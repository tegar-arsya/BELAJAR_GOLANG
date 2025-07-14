package helpers

import (
	"fmt"
	"time"
)

func GenerateFileName(original string) string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("img_%d_%s", timestamp, original)
}
