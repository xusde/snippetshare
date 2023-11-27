/**
 * Testing file for templates.go
 */
package main

import (
	"testing"
	"time"

	"github.com/xusde/snippetshare/internal/assert"
)

func TestHumanDate(t *testing.T) {

	// Create a slice of anonymous structs containing the test case name,
	// input to our humanDate() function (the tm field), and expected output
	// (the want field).
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2021, 03, 20, 10, 10, 0, 0, time.UTC),
			want: "Mar 20 2021 at 10:10",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2021, 03, 20, 10, 10, 0, 0, time.FixedZone("CET", 60*60)),
			want: "Mar 20 2021 at 09:10",
		},
	}

	// Loop over the test cases
	for _, tt := range tests {
		// Use the t.Run() function to run a sub-test for each test case.
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)
			// check if the returned humanDate is equal to the expected humanDate
			assert.Equal(t, hd, tt.want)
		})
	}
}
