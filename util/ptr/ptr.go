package ptr

import "time"

func Str(s string) *string                    { return &s }
func Int(i int) *int                          { return &i }
func Duration(d time.Duration) *time.Duration { return &d }
func Time(t time.Time) *time.Time             { return &t }
func Bool(b bool) *bool                       { return &b }
