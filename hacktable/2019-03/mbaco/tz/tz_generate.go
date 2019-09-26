package tz

//go:generate rm -fr zoneinfo
//go:generate unzip -q $GOROOT/lib/time/zoneinfo.zip -d zoneinfo/
//go:generate go run scripts/generate_timezone_data.go -o zoneinfo_windows.go
//go:generate rm -fr zoneinfo
