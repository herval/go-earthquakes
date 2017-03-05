package earthquakes

import (
	"encoding/json"
	"fmt"
	"github.com/dropbox/godropbox/errors"
	"net/http"
	"time"
)

type Timeframe int

const (
	AllHour              Timeframe = iota
	SignificantHour
	AllDay
	SignificantDay
	AllSevenDays
	SignificantSevenDays
	AllMonth
	SignificantMonth
)

// Format defined on https://earthquake.usgs.gov/earthquakes/feed/v1.0/geojson.php
type Earthquake struct {
	Magnitude float32 `json:"mag"`
	Alert     string `json:"alert"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	Tsunami   int `json:"tsunami"`
	Time      int `json:"time"`
}

func Feed(timeframe Timeframe) ([]Earthquake, error) {
	conf := summaryConfigs[timeframe]
	if conf == nil {
		return nil, errors.New("Unsupported timeframe")
	}

	// getting
	url := fmt.Sprintf("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/%s", conf.endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parsing
	var record featureCollection
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		return nil, err
	}

	// picking
	quakes := make([]Earthquake, 0)
	for _, f := range record.Features {
		quakes = append(quakes, f.Properties)
	}
	return quakes, err
}

type attrs struct {
	refreshInterval time.Duration
	endpoint        string
}

var summaryConfigs = map[Timeframe]*attrs{
	AllHour: &attrs{
		endpoint: "significant_hour.geojson",
	},
	SignificantHour: &attrs{
		endpoint: "significant_hour.geojson",
	},
	AllDay: &attrs{
		endpoint: "all_day.geojson",
	},
	SignificantDay: &attrs{
		endpoint: "significant_day.geojson",
	},
	AllSevenDays: &attrs{
		endpoint: "all_week.geojson",
	},
	SignificantSevenDays: &attrs{
		endpoint: "significant_week.geojson",
	},
	AllMonth: &attrs{
		endpoint: "all_month.geojson",
	},
	SignificantMonth: &attrs{
		endpoint: "significant_month.geojson",
	},
}

type featureCollection struct {
	Features []feature `json:"features"`
}

type feature struct {
	Properties Earthquake `json:"properties"`
}
