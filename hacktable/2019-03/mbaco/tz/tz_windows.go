package tz

import "time"

func LoadLocation(name string) (*time.Location, error) {
	data, err := zoneData(name)
	if err != nil {
		return nil, err
	}

	return time.LoadLocationFromTZData(name, data)
}
