package main

import "testing"

func TestValidateMovie(t *testing.T) {
	tests := []struct {
		name    string
		movie   Movie
		wantErr string // "" means no error expected
	}{
		// new(value) returns a pointer to that value (Go 1.26+), handy for the *int / *float64 fields.
		{"valid movie", Movie{Title: "Whiplash", ReleaseYear: new(2014), Rating: new(8.5)}, ""},
		{"rating of 10 is allowed", Movie{Title: "Whiplash", Rating: new(10.0)}, ""},
		{"empty title", Movie{Title: ""}, "title is required"},
		{"rating too high", Movie{Title: "Whiplash", Rating: new(11.0)}, "rating must be between 0 and 10"},
		{"negative rating", Movie{Title: "Whiplash", Rating: new(-1.0)}, "rating must be between 0 and 10"},
		{"year too early", Movie{Title: "Whiplash", ReleaseYear: new(1500)}, "release_year must be 1888 or later"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateMovie(tt.movie)

			got := ""
			if err != nil {
				got = err.Error()
			}

			if got != tt.wantErr {
				t.Errorf("validateMovie() error = %q, want %q", got, tt.wantErr)
			}
		})
	}
}
