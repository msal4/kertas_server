package ptr

import "github.com/msal4/hassah_school_server/ent/schema"

// Str takes a string and returns a pointer to that string.
func Str(s string) *string { return &s }

// Int takes an integer and returns a pointer to that integer.
func Int(i int) *int { return &i }

// Bool takes a boolean and returns a pointer to that boolean.
func Bool(b bool) *bool { return &b }

// Status takes an entity status and returns a pointer to that status.
func Status(s schema.Status) *schema.Status { return &s }
