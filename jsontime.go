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

	// 首先看是否是数字，表示毫秒数, 1milli=1000,000nano
	if p, err := strconv.ParseInt(v, 10, 64); err == nil {
		*t = Time(time.Unix(0, p*1000000))
		return nil
	}

	v = strings.ReplaceAll(v, ",", ".")

	for _, f := range []string{
		"2006-01-02T15:04:05.000000Z",
		"2006-01-02T15:04:05.000",
		"2006-01-02 15:04:05.000",
	} {
		if tt, err := time.Parse(f, v); err == nil {
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
