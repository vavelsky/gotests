package selects

import "testing"

func TestRacer(t *testing.T) {
	slowUrl := "http://facebook.com"
	fastUrl := "http://www.quii.co.uk"

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
