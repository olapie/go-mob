package mobile

import (
	"code.olapie.com/types"
)

type Point types.Point

func NewPoint() *Point {
	return new(Point)
}

type Place types.Place

func NewPlace() *Place {
	return new(Place)
}

func (p *Place) SetCoordinate(c *Point) {
	p.Coordinate = (*types.Point)(c)
}

func (p *Place) GetCoordinate() *Point {
	return (*Point)(p.Coordinate)
}

type Country types.Country

func GetCountryList() *CountryList {
	l := GetCountries()
	ll := new(CountryList)
	ll.Elements = make([]*Country, len(l))
	for i, c := range l {
		ll.Elements[i] = (*Country)(c)
	}
	return ll
}
