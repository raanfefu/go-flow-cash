package finances

import (
	"strings"
	"time"
)

func DateToInt64(date time.Time) int64 {
	num := float64(0)
	return int64(num + date.Sub(MakeDate(1900, 1, 1)).Hours()/24)
}
func DateToFloat64(date time.Time) float64 {
	num := float64(0)
	return float64(num + date.Sub(MakeDate(1900, 1, 1)).Hours()/24)
}

func MakeDate(year int, month int, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, &time.Location{})
}

func (c *CustomDate) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*c = CivilTime(t) //set result using the pointer
	return nil
}
