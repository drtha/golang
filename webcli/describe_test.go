package webcli

import "testing"

func TestFormatDate (t *testing.T) {
	in := "2011-11-11T11:11:11+01:00"
	want := "11.11.2011 um 11:11"

	out := formatDate(in)
	if out != want {
		t.Errorf("Date was wrong, got: %s, want: %s.", out, want)
	}
}

