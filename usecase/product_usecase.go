package usecase

import (
	"github.com/alexander-pastana/go-api-lab/model"
	"github.com/alexander-pastana/go-api-lab/repository"
	"errors"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()

}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {

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

func (pu *ProductUsecase) UpdateProduct(product model.Product) (model.Product, error) {
	//Checa se o id existe para poder atualizar
	_, err := pu.GetProductById(product.ID)
	//tratativa de erro
	if err != nil {
		return model.Product{}, errors.New("Produto não encontrado")
	}
	//func de atualizar produto
	err = pu.repository.UpdateProduct(product)
	//tratativa de erro
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}
