// Copyright (C) 2017 ScyllaDB

package format

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// StringByteCount returns string representation of the byte count with proper
// unit.
func StringByteCount(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	// No decimals by default, two decimals for GiB and three for more than
	// that.
	format := "%.0f%ciB"
	if exp == 2 {
		format = "%.2f%ciB"
	} else if exp > 2 {
		format = "%.3f%ciB"
	}
	return fmt.Sprintf(format, float64(b)/float64(div), "KMGTPE"[exp])
}

var (
	byteCountRe          = regexp.MustCompile(`([0-9]+(?:\.[0-9]+)?)(B|[KMGTPE]iB)`)
	byteCountReValueIdx  = 1
	byteCountReSuffixIdx = 2
)

// ParseByteCount returns byte count parsed from input string.
// This is opposite of StringByteCount function.
func ParseByteCount(s string) (int64, error) {
	const unit = 1024
	var exps = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
	parts := byteCountRe.FindStringSubmatch(s)
	if len(parts) != 3 {
		return 0, errors.Errorf("invalid byte size string: %q; it must be real number with unit suffix: %s", s, strings.Join(exps, ","))
	}

	v, err := strconv.ParseFloat(parts[byteCountReValueIdx], 64)
	if err != nil {
		return 0, errors.Wrapf(err, "parsing value for byte size string: %s", s)
	}

	pow := 0
	for i, e := range exps {
		if e == parts[byteCountReSuffixIdx] {
			pow = i
		}
	}

	mul := math.Pow(unit, float64(pow))

	return int64(v * mul), nil
}
