package mobile

import (
	_ "embed"
	"encoding/json"
	"strings"
	"sync"

	"code.olapie.com/log"
	"github.com/nyaruka/phonenumbers"
)

//go:embed resource/countries.json
var countriesJSONString string

func GetCountries() []*Country {
	loadCountries()
	return countries
}

func GetCountryByCallingCode(code int64) *Country {
	loadCountries()
	return codeToCountry[code]
}

func GetCallingCodeByRegion(regionCode string) int {
	return phonenumbers.GetCountryCodeForRegion(strings.ToUpper(regionCode))
}

var countries []*Country
var codeToCountry map[int64]*Country
var loadCountriesOnce sync.Once

func loadCountries() {
	loadCountriesOnce.Do(func() {
		if err := json.Unmarshal([]byte(countriesJSONString), &countries); err != nil {
			log.Error(err)
			countries = []*Country{}
			return
		}

		codeToCountry = make(map[int64]*Country, len(countries))
		for _, c := range countries {
			codeToCountry[c.CallingCode] = c
		}
	})
}
