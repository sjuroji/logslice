package timeparser_test

import (
	"testing"
	"time"

	"github.com/yourorg/logslice/internal/timeparser"
)

func TestParse_SupportedFormats(t *testing.T) {
	p := timeparser.New(time.UTC)

	cases := []struct {
		input string
		wantYear int
		wantMonth time.Month
		wantDay int
	}{
		{"2024-03-15T10:20:30Z", 2024, time.March, 15},
		{"2024-03-15T10:20:30.123456789Z", 2024, time.March, 15},
		{"2024-03-15 10:20:30", 2024, time.March, 15},
		{"2024-03-15 10:20:30.999", 2024, time.March, 15},
		{"2024-03-15T10:20:30", 2024, time.March, 15},
		{"2024-03-15", 2024, time.March, 15},
		{"15/Mar/2024:10:20:30 +0000", 2024, time.March, 15},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := p.Parse(tc.input)
			if err != nil {
				t.Fatalf("Parse(%q) unexpected error: %v", tc.input, err)
			}
			if got.Year() != tc.wantYear || got.Month() != tc.wantMonth || got.Day() != tc.wantDay {
				t.Errorf("Parse(%q) = %v, want %d-%02d-%02d",
					tc.input, got, tc.wantYear, tc.wantMonth, tc.wantDay)
			}
		})
	}
}

func TestParse_UnknownFormat(t *testing.T) {
	p := timeparser.New(nil)
	_, err := p.Parse("not-a-time")
	if err == nil {
		t.Fatal("expected error for unknown format, got nil")
	}
}

func TestParseWithFormat(t *testing.T) {
	p := timeparser.New(time.UTC)
	layout := "02 Jan 2006 15:04"
	input := "15 Mar 2024 10:20"

	got, err := p.ParseWithFormat(layout, input)
	if err != nil {
		t.Fatalf("ParseWithFormat unexpected error: %v", err)
	}
	if got.Day() != 15 || got.Month() != time.March || got.Year() != 2024 {
		t.Errorf("unexpected result: %v", got)
	}
}

func TestNew_NilLocation(t *testing.T) {
	p := timeparser.New(nil)
	got, err := p.Parse("2024-01-02")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Location() != time.UTC {
		t.Errorf("expected UTC location, got %v", got.Location())
	}
}
