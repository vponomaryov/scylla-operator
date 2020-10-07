// Copyright (C) 2017 ScyllaDB

package format

import (
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/scylladb/scylla-mgmt-commons/timeutc"
)

const (
	nowSafety = 30 * time.Second
)

// ParseStartDate parses the supplied string as a time.Time.
func ParseStartDate(value string) (time.Time, error) {
	now := timeutc.Now()

	if value == "now" {
		return now.Add(nowSafety), nil
	}

	if strings.HasPrefix(value, "now") {
		d, err := time.ParseDuration(value[3:])
		if err != nil {
			return time.Time{}, err
		}
		if d < 0 {
			return time.Time{}, errors.New("start date cannot be in the past")
		}
		if d < nowSafety {
			return time.Time{}, errors.Errorf("start date must be at least in %s", nowSafety)
		}
		return now.Add(d), nil
	}

	// No more heuristics, assume the user passed a date formatted string
	t, err := timeutc.Parse(time.RFC3339, value)
	if err != nil {
		return t, err
	}
	if t.Before(now) {
		return time.Time{}, errors.New("start date cannot be in the past")
	}
	if t.Before(now.Add(nowSafety)) {
		return time.Time{}, errors.Errorf("start date must be at least in %s", nowSafety)
	}
	return t, nil
}
