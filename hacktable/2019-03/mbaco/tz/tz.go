// +build !windows

package tz

import "time"

func LoadLocation(name string) (*time.Location, error) {
	return time.LoadLocation(name)
}
