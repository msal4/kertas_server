package ptr

// Str takes a string and returns a pointer to that string.
func Str(s string) *string { return &s }

// Int takes an integer and returns a pointer to that integer.
func Int(i int) *int { return &i }

// Bool takes a boolean and returns a pointer to that boolean.
func Bool(b bool) *bool { return &b }
