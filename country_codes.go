package bbcweather

import (
	_ "embed"
	"encoding/json"
	"sync"
)

var (
	//go:embed country_codes.json
	countryCodesData string

	once         sync.Once
	countryCodes map[string]string
)

type rawCodeData []struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// GetCountryFromCode useds an embedded list of country codes to find the corresponding country name
func GetCountryFromCode(code string) (string, bool) {
	once.Do(func() {
		var rawData rawCodeData
		err := json.Unmarshal([]byte(countryCodesData), &rawData)
		if err != nil {
			panic(err)
		}
		countryCodes = make(map[string]string)
		for _, entry := range rawData {
			countryCodes[entry.Code] = entry.Name
		}
	})
	result, ok := countryCodes[code]
	return result, ok
}
