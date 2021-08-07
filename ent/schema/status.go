package schema

type Status string

// StatusActive is the default value of the Status enum.
const DefaultStatus = StatusActive

// Status values.
const (
	StatusDeleted  Status = "deleted"
	StatusDisabled Status = "disabled"
	StatusActive   Status = "active"
)

func (s Status) String() string {
	return string(s)
}

func (Status) Values() []string {
	return []string{StatusDeleted.String(), StatusDisabled.String(), StatusActive.String()}
}
