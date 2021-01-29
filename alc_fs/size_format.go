package alc_fs

import (
	"strconv"
	"strings"
)

const (
	FileSizeB = 1 << (10 * iota)
	FileSizeKB
	FileSizeMB
	FileSizeGB
	FileSizeTB
	FileSizePB
	FileSizeEB
)

func SizeFormat(bytes int64) string {
	unit := ""
	value := float64(bytes)

	switch {
	case bytes >= FileSizeEB:
		unit = "E"
		value = value / FileSizeEB
	case bytes >= FileSizePB:
		unit = "P"
		value = value / FileSizePB
	case bytes >= FileSizeTB:
		unit = "T"
		value = value / FileSizeTB
	case bytes >= FileSizeGB:
		unit = "G"
		value = value / FileSizeGB
	case bytes >= FileSizeMB:
		unit = "M"
		value = value / FileSizeMB
	case bytes >= FileSizeKB:
		unit = "K"
		value = value / FileSizeKB
	case bytes >= FileSizeB:
		unit = "B"
	case bytes == 0:
		return "0B"
	}

	result := strconv.FormatFloat(value, 'f', 1, 64)
	result = strings.TrimSuffix(result, ".0")
	return result + unit
}
