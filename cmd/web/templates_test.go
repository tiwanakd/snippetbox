package main

import (
	"testing"
	"time"

	"github.com/tiwanakd/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 9, 25, 17, 30, 0, 0, time.UTC),
			want: "25 Sep 2024 at 17:30",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, 9, 25, 17, 30, 0, 0, time.FixedZone("PDT", 7*60*60)),
			want: "25 Sep 2024 at 10:30",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)
			assert.Equal(t, hd, tt.want)
		})
	}
}
