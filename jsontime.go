package jsontime

import (
	"errors"
	perrors "github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

// ErrUnknownTimeFormat defines the error type for unknown time format.
var ErrUnknownTimeFormat = errors.New("unkown errors time format")

// Time defines a time.Time that can be used in struct tag for JSON unmarshalling.
type Time time.Time

// UnmarshalJSON unmarshals bytes to Time.
func (t *Time) UnmarshalJSON(b []byte) error {
	v, _ := TryUnQuoted(string(b))
	if v == "" {
		return nil
	}

	// 首先看是否是数字，表示毫秒数或者纳秒数
	if p, err := strconv.ParseInt(v, 10, 64); err == nil {
		*t = Time(ParseTime(p))
		return nil
	}

	v = strings.ReplaceAll(v, ",", ".")
	v = strings.ReplaceAll(v, "T", " ")
	v = strings.ReplaceAll(v, "-", "")
	v = strings.ReplaceAll(v, ":", "")
	v = strings.TrimSuffix(v, "Z")

	for _, f := range []string{"20060102 150405.000000", "20060102 150405.000"} {
		if tt, err := time.ParseInLocation(f, v, time.Local); err == nil {
			*t = Time(tt)
			return nil
		}
	}

	return perrors.Wrapf(ErrUnknownTimeFormat, "value %s has unknown time format"+v)
}

// TryUnQuoted tries to unquote string.
func TryUnQuoted(v string) (string, bool) {
	vlen := len(v)
	yes := vlen >= 2 && v[0] == '"' && v[vlen-1] == '"'
	if !yes {
		return v, false
	}

	return v[1 : vlen-1], true
}

// ParseTime tries to parse a int64 value to a time in this year by seconds, milliseconds,or  nanoseconds.
func ParseTime(v int64) time.Time {
	t := time.Now()
	yearStart := time.Date(t.Year()-1, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	yearEnd := time.Date(t.Year()+1, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	if yearStart.Unix() <= v && v < yearEnd.Unix() {
		return time.Unix(v, 0) // seconds range
	}

	if yearStart.Unix()*1000 <= v && v < yearEnd.Unix()*1000 {
		return time.Unix(0, v*1000000) // milliseconds range, 1 millis = 1000,000 nanos
	}

	if yearStart.UnixNano() <= v && v < yearEnd.UnixNano() {
		return time.Unix(0, v) // nanoseconds range
	}

	return time.Unix(v, 0)
}
