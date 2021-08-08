package schema

import (
	"fmt"
	"io"
)

type Status string

// StatusActive is the default value of the Status enum.
const DefaultStatus = StatusActive

// Status values.
const (
	StatusDeleted  Status = "DELETED"
	StatusDisabled Status = "DISABLED"
	StatusActive   Status = "ACTIVE"
)

func (s Status) String() string {
	return string(s)
}

func (Status) Values() []string {
	return []string{StatusDeleted.String(), StatusDisabled.String(), StatusActive.String()}
}

func (s Status) IsValid() bool {
	switch s {
	case StatusDeleted, StatusDisabled, StatusActive:
		return true
	default:
		return false
	}
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (s *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("status must be a string")
	}

	status := Status(str)
	if !status.IsValid() {
		return fmt.Errorf("%q is not a valid status", v)
	}

	*s = status

	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (s Status) MarshalGQL(w io.Writer) {
	w.Write([]byte(s))
}
