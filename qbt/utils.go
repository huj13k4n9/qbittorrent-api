package qbt

import (
	"strconv"
	"time"
)

type Time time.Time

const URLPattern = "%s/api/v2/%s"
const Version = "v0.1"

func (t *Time) UnmarshalJSON(bytes []byte) error {
	timestamp, err := strconv.Atoi(string(bytes))
	if err != nil {
		return err
	}
	*t = Time(time.Unix(int64(timestamp), 0))
	return nil
}
