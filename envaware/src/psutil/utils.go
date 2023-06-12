package psutil

import "fmt"

const (
	KIBI uint64 = 1 << 10
	MEBI uint64 = 1 << 20
	GIBI uint64 = 1 << 30
	TEBI uint64 = 1 << 40
	PEBI uint64 = 1 << 50
	EXBI uint64 = 1 << 60
)

// formatBytes into a rounded string representation using IEC standard
func FormatBytes(bytes uint64) string {
	if bytes == 1 { // bytes
		return fmt.Sprintf("%d byte", bytes)
	} else if bytes < KIBI { // bytes
		return fmt.Sprintf("%d bytes", bytes)
	} else if bytes < MEBI { // KiB
		return formatUnits(bytes, KIBI, "KiB")
	} else if bytes < GIBI { // MiB
		return formatUnits(bytes, MEBI, "MiB")
	} else if bytes < TEBI { // GiB
		return formatUnits(bytes, GIBI, "GiB")
	} else if bytes < PEBI { // TiB
		return formatUnits(bytes, TEBI, "TiB")
	} else if bytes < EXBI { // PiB
		return formatUnits(bytes, PEBI, "PiB")
	} else { // EiB
		return formatUnits(bytes, EXBI, "EiB")
	}
}

// formatUnits as exact integer or fractional decimal based on the prefix,appending the appropriate units
func formatUnits(value uint64, prefix uint64, unit string) string {
	if value%prefix == 0 {
		return fmt.Sprintf("%d %s", value/prefix, unit)
	}
	return fmt.Sprintf("%.1f %s", float64(value)/float64(prefix), unit)
}
