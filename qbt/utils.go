package qbt

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"
)

type Time time.Time

const URLPattern = "%s/api/v2/%s"
const Version = "v0.1"

var ErrBadResponse = errors.New("received bad response")
var ErrUnknownType = errors.New("unknown type")
var ErrUnauthenticated = errors.New("unauthenticated request")

func WriteFile(path string, content []byte, overwrite bool) error {
	flags := os.O_CREATE | os.O_WRONLY
	if !overwrite {
		flags |= os.O_EXCL
	} else {
		flags |= os.O_TRUNC
	}

	file, err := os.OpenFile(path, flags, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		return err
	}
	return nil
}

func MapToStruct[T interface{}](mapData map[string]any, structData *T) error {
	bytes, err := json.Marshal(mapData)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, structData)
	if err != nil {
		return err
	}
	return nil
}

func (t *Time) UnmarshalJSON(bytes []byte) error {
	timestamp, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		return err
	}
	*t = Time(time.Unix(timestamp, 0))
	return nil
}
