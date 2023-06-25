package usecase

import (
	"errors"
	"fmt"
	"realimage/domain"
	"realimage/repository"
)

var cRepo *repository.ContributorRepository
var repo *repository.LocationRepository

func init() {
	cRepo = repository.NewContributorRepository()
	repo = repository.NewLocationRepository()
}

func CreateContributor(contributorId int) error {
	contributor := domain.Contributor{
		ID: contributorId,
	}
	err := cRepo.CreateContributor(&contributor)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func SetExludedContributorLocations(start, end, contributorId int) error {
	locations, err := repo.GetLocationsByRange(start, end)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(locations)
	}

	contributor, err := cRepo.GetContributorById(contributorId)
	if err != nil {
		return err
	}

	for _, location := range locations {
		contributor.ExcludedLocations = append(contributor.ExcludedLocations, location)
	}
	fmt.Println(contributor)
	return nil

}

func SetIncludedContributorLocations(start, end, contributorId int) error {
	locations, err := repo.GetLocationsByRange(start, end)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(locations)
	}

	contributor, err := cRepo.GetContributorById(contributorId)
	if err != nil {
		return err
	}
	includedCountries := make(map[string]bool)

	for _, location := range locations {
		if !includedCountries[location.CountryName] {
			contributor.IncludedLocations = append(contributor.IncludedLocations, location.CountryName)
			includedCountries[location.CountryName] = true
		}

	}
	fmt.Println(contributor)
	return nil
}

func CheckContributorPermission(contributorId int, country, province, city string) (error, bool) {
	var permission bool = false
	contributor, err := cRepo.GetContributorById(contributorId)
	if err != nil {
		return err, permission
	}
	for _, loc := range contributor.IncludedLocations {
		if loc == country {
			permission = true
			break
		}
	}
	if permission == false {
		return errors.New("Permission denied for the country"), permission
	}
	for _, loc := range contributor.ExcludedLocations {
		if loc.CountryName == country && loc.ProvinceName == province && loc.CityName == city {
			permission = false
		}
	}
	return nil, permission

}
