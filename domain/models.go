package domain

type Location struct {
	CityCode     string
	ProvinceCode string
	CountryCode  string
	CityName     string
	ProvinceName string
	CountryName  string
}

type Contributor struct {
	ID                int
	IncludedLocations []string
	ExcludedLocations []Location
}
