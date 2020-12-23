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
	if numBytes > TB {
		return fmt.Sprintf("%.1fTB", numBytesF/TB)
	}
	if numBytes > GB {
		return fmt.Sprintf("%.1fGB", numBytesF/GB)
	}
	if numBytes > MB {
		return fmt.Sprintf("%.1fMB", numBytesF/MB)
	}
	if numBytes > KB {
		return fmt.Sprintf("%.1fKB", numBytesF/KB)
	}
	// remaining is just bytes
	return fmt.Sprintf("%dB", numBytes)
}
