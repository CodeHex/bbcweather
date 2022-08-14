# BBC Weather API tools
[![Go Reference](https://pkg.go.dev/badge/github.com/codehex/bbcweather.svg)](https://pkg.go.dev/github.com/codehex/bbcweather)

Retrieves the current and forecasted weather from the BBC Weather APIs

# wx
`wx` is a command line tool for displaying the current weather from your nearest observation station and the forecasted weather for the next 14 days. Uses the BBC weather APIs for the source of information.
```zsh
go install github.com/codehex/bbcweather/cmd/wx@latest
```
Locations can be in the form of
- The name of the location, city or region
- The first part of the post code
- Longitude and latitude
- The location ID (see [Geonames](https://www.geonames.org/) dataset for more details on location IDs, data dump is [here](http://download.geonames.org/export/dump/))
e.g.

```shell
wx Edinburgh Airport     # By specific location
wx New York              # By city or town
wx Lake District         # By region
wx Switzerland           # By country
wx SW19                  # By postcode (e.g. Wimbeldon)
wx 48.8566,2.3522        # By longitude and latitude (e.g. Paris)
wx 2993458               # By location ID (e.g. Monaco - https://www.geonames.org/2993458)
```

A preset location can be used by setting the `WX_QUERY` env variable
```zsh
export WX_QUERY=Paris
wx
```
Here's an example of a weather report
<p align="center">
<img width="779" alt="wx Output" src="https://user-images.githubusercontent.com/3924123/184547884-53e486fe-6deb-48fe-ad93-dfd281204625.png">
</p>


