package str

import "time"

func ToTime(str, layout string) time.Time {
	time, _ := time.Parse(layout, str)
	return time
}

func ToNilTime(str *string, layout string) *time.Time {
	if str != nil {
		time, _ := time.Parse(layout, *str)
		return &time
	}
	return nil
}
