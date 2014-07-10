package conekta

import (
	"encoding/json"
	"strconv"
	"time"
)

var (
	layout = "2006-01-02"
)

type timestamp struct {
	time.Time
}

func NewTimestamp(timeStr string) *timestamp {
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		panic("Cannot parse time layout")
	}
	return &timestamp{t}
}

func (t timestamp) String() string {
	return t.Time.String()
}

func (ts *timestamp) UnmarshalJSON(b []byte) error {
	result, err := strconv.ParseInt(string(b), 10, 64)
	if err == nil {
		(*ts).Time = time.Unix(result, 0)
	} else {
		(*ts).Time, err = time.Parse(`"`+time.RFC3339+`"`, string(b))
	}
	return err
}

func (ts *timestamp) MarshalJSON() ([]byte, error) {
	str := (*ts).Time.Format(layout)
	return json.Marshal(str)
}
