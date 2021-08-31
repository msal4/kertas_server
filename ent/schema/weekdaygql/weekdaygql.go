package weekdaygql

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalWeekday(t time.Weekday) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Itoa(int(t)))
	})
}

func UnmarshalWeekday(v interface{}) (time.Weekday, error) {
	switch v := v.(type) {
	case int:
		return time.Weekday(v), nil
	case string:
		i, err := strconv.Atoi(v)
		return time.Weekday(i), err
	case json.Number:
		i, err := v.Int64()
		if err != nil {
			return 0, err
		}
		return time.Weekday(i), nil
	default:
		return 0, fmt.Errorf("invalid type %T, expected string, int or json number", v)
	}
}
