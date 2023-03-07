package types

import "time"

const layout = `"2006-01-02T15:04:05.000+0000"`

type Time time.Time

func (t *Time) UnmarshalJSON(bytes []byte) error {
	parsed, err := time.Parse(layout, string(bytes))
	if err != nil {
		return err
	}
	*t = Time(parsed)
	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	str := time.Time(*t).Format(layout)
	return []byte(str), nil
}
