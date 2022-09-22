package service

import (
	"ASSIGNMENT2/models"
	brandRepo "ASSIGNMENT2/repository"
)

func GetBrand(brandID uint64) (brand models.Brand, err error) {
	brand, err = brandRepo.GetBrand(brandID)
	return
}

func CreateBrand(newBrand models.Brand) (models.Brand, error) {
	created, err := brandRepo.CreateBrand(newBrand)
	return created, err
}

func UpdateBrand(newBrand models.Brand) (models.Brand, int64, error) {
	updated, affected, err := brandRepo.UpdateBrand(newBrand)
	return updated, affected, err
}
