package product

import (
	"crud-golang/config/database"

	"crud-golang/config"
)

type Repository interface {
	Create(product Product) (Product, error)
	FindByID(id int) (Product, error)
	FindAll() ([]Product, error)
	Update(product Product) (Product, error)
	Delete(product Product) (Product, error)
}

type repository struct {
	cfg config.Config
	DB  database.GormDatabase
}

func NewRepository(cfg config.Config) *repository {
	return &repository{cfg: cfg, DB: cfg.DB()}
}

func (r *repository) Create(product Product) (Product, error) {
	if err := r.DB.Master().Create(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindAll() ([]Product, error) {
	var product []Product
	if err := r.DB.Master().Find(&product).Error; err != nil {
		return product, err
	}
	return product, nil

}

func (r *repository) FindByID(id int) (Product, error) {
	var product Product
	if err := r.DB.Master().Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Update(product Product) (Product, error) {
	if err := r.DB.Master().Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Delete(product Product) (Product, error) {
	if err := r.DB.Master().Delete(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}
