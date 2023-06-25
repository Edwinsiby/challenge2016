package repository

import (
	"encoding/csv"
	"fmt"
	"os"
	"realimage/domain"
)

type LocationRepository struct {
	filePath string
}

func NewLocationRepository() *LocationRepository {
	return &LocationRepository{
		filePath: "cities.csv",
	}
}

type ContributorRepository struct {
	contributors map[int]*domain.Contributor
}

func NewContributorRepository() *ContributorRepository {
	return &ContributorRepository{
		contributors: make(map[int]*domain.Contributor),
	}
}

func (r *ContributorRepository) CreateContributor(contributer *domain.Contributor) error {
	r.contributors[contributer.ID] = contributer
	fmt.Println(r.contributors)
	return nil
}

func (r *ContributorRepository) GetContributorById(id int) (*domain.Contributor, error) {
	fmt.Println(r.contributors)
	contributor, ok := r.contributors[id]
	if !ok {
		return nil, fmt.Errorf("contributor not found")
	}
	return contributor, nil
}

func (r *LocationRepository) GetLocationsByLimit(limit int) ([]domain.Location, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	var records [][]string
	// Read the specified number of records from the CSV file
	for i := 0; i < limit; i++ {
		record, err := reader.Read()
		if err != nil {
			break // Reached end of file or encountered an error
		}
		records = append(records, record)
	}

	// Process records and populate locations
	var locations []domain.Location
	for _, record := range records {
		location := domain.Location{
			CityCode:     record[0],
			ProvinceCode: record[1],
			CountryCode:  record[2],
			CityName:     record[3],
			ProvinceName: record[4],
			CountryName:  record[5],
		}

		locations = append(locations, location)
	}

	return locations, nil
}

func (r *LocationRepository) GetLocationsByRange(start, end int) ([]domain.Location, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var records [][]string
	for i := 0; i <= end; i++ {
		record, err := reader.Read()
		if err != nil {
			break
		}
		if i > start {
			records = append(records, record)
		}

	}
	var locations []domain.Location
	for _, record := range records {
		location := domain.Location{
			CityCode:     record[0],
			ProvinceCode: record[1],
			CountryCode:  record[2],
			CityName:     record[3],
			ProvinceName: record[4],
			CountryName:  record[5],
		}

		locations = append(locations, location)
	}

	return locations, nil
}
