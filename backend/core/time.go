package core

import "time"

var JST = time.FixedZone("Asia/Tokyo", 9*60*60)

const (
	LayoutISO8601   = "2006-01-02T15:04:05Z07:00"
	LayoutDateTime  = "2006-01-02 15:04:05"
	LayoutDateOnly  = "2006-01-02"
	LayoutDateSlash = "2006/01/02"
	LayoutTimeOnly  = "15:04:05"
)

func NowJST() time.Time {
	return time.Now().In(JST)
}

func ToJST(t time.Time) time.Time {
	return t.In(JST)
}

func ParseJST(layout, value string) (time.Time, error) {
	t, err := time.ParseInLocation(layout, value, JST)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func FormatJST(t time.Time, layout string) string {
	return t.In(JST).Format(layout)
}

func FormatWithLayout(t time.Time, layout string) string {
	return t.Format(layout)
}
