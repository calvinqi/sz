package helpers

import (
	"fmt"
)

// constants for magnitude conversions
const (
	TB = 1000000000000
	GB = 1000000000
	MB = 1000000
	KB = 1000
)

// ReadableBytes convert byte num to readable string. E.g. 1240 -> 1.2KB
func ReadableBytes(numBytes int64) string {
	numBytesF := float64(numBytes)

	amt := numBytesF
	unitStr := "B"
	if numBytes > TB {
		amt = numBytesF / TB
		unitStr = "T"
	} else if numBytes > GB {
		amt = numBytesF / GB
		unitStr = "G"
	} else if numBytes > MB {
		amt = numBytesF / MB
		unitStr = "M"
	} else if numBytes > KB {
		amt = numBytesF / KB
		unitStr = "K"
	}
	// remaining case is just bytes
	if amt < 10 {
		return fmt.Sprintf("%.1f%s", amt, unitStr)
	}
	return fmt.Sprintf("%d%s", int(amt), unitStr)
}
