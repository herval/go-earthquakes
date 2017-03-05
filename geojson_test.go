package earthquakes_test

import (
	"fmt"
	"github.com/herval/go-earthquakes"
	"testing"
)

func TestFetch(t *testing.T) {
	quakes, err := earthquakes.Feed(earthquakes.AllDay)
	if err != nil {
		t.Fatal(err)
	}

	for _, quake := range quakes {
		fmt.Printf(" %+v\n", quake)
	}
}
