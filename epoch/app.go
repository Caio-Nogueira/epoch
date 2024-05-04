package epoch

import (
	"fmt"
	"time"
)

func location() *time.Location {
	loc, _ := time.LoadLocation("Europe/Lisbon")
	return loc
}

func Epoch(date string) (int, error) {
	const layout = "2006-01-02T15:04"
	t, err := time.Parse(layout, date)
	if err != nil {
		return -1, err
	}

	return int(t.Unix()), nil
}

func Date(epoch int) string {
	loc := location()
	t := time.Unix(int64(epoch), 0).In(loc)
	return t.Format(time.RFC3339)
}

func CurrentTime() string {
	loc := location()
	msg := "{Current time (UNIX): %v, \nCurrent time (RFC3339): %v}"
	return fmt.Sprintf(msg, time.Now().In(loc).Unix(), time.Now().In(loc).Format(time.RFC3339))
}

