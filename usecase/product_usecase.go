package usecase

import (
	"errors"
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CretaeProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUsecase) DeleteProduct(id_product int) error {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return err
	}

	if product == nil {
		return errors.New("product not found")
	}

	err = pu.repository.DeleteProduct(id_product)
	if err != nil {
		return err
	}

	return nil
}

func (pu *ProductUsecase) UpdateProduct(product model.Product) (*model.Product, error) {
	existingProduct, err := pu.repository.GetProductById(product.ID)
	if err != nil {
		return nil, err
	}

	if existingProduct == nil {
		return nil, errors.New("product not found")
	}

	updatedProduct, err := pu.repository.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}
