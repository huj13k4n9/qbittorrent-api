package qbt

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Time time.Time

const URLPattern = "%s/api/v2/%s"
const Version = "v0.1"

var ErrBadResponse = errors.New("received bad response")
var ErrUnauthenticated = errors.New("unauthenticated request")

func (t *Time) UnmarshalJSON(bytes []byte) error {
	timestamp, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		return err
	}
	*t = Time(time.Unix(timestamp, 0))
	return nil
}
