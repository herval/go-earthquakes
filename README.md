# Go-Earthquakes

A small Go wrapper for the [USGS Earthquake Data APIs](https://earthquake.usgs.gov).

Started this inspired by Moon Ribas' [Waiting for Earthquakes performance](https://www.youtube.com/watch?v=1Un4MFR-vNI) - using it to building my own little Earthquake "sensor" to put my Raspberry Pi Zero W to use :-) (a bit of an "heavy handed" hardware, I know - more to come soon!)

Current version supports only the [GeoJSON](https://earthquake.usgs.gov/earthquakes/feed/v1.0/geojson.php) feeds (hourly, daily, weekly and monthly).

# TODO

- Add support for the "geometry" field (event coordinates)
- Support general [Query](https://earthquake.usgs.gov/fdsnws/event/1/) API