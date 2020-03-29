package services

import (
	"fmt"
	"test-produce/datamodels"
	"test-produce/repositories"
)

// 服务管理员接口
type IProductService interface {
	GetProductByID(int64) (*datamodels.Product,error)
	GetAllProduct()([]*datamodels.Product,error)
	DeleteProductByID(int64)bool
	InsertProduct(product *datamodels.Product)(int64,error)
	UpdateProduct(product *datamodels.Product)error
}

// 服务管理员
type ProductService struct {
	productRepository repositories.IProduct
}

// 初始化服务管理员函数
func NewProductService(repository repositories.IProduct)IProductService   {
	return &ProductService{repository}
}

// 通过ID获取商品函数
func (p *ProductService)GetProductByID(productID int64) (*datamodels.Product,error)  {
	return p.productRepository.SelectByKey(productID)
}

// 获取所有商品函数
func (p *ProductService) GetAllProduct() ([]*datamodels.Product, error) {
	fmt.Println("all: step1")
	return p.productRepository.SelectAll()
}

// 删除函数
func (p *ProductService)DeleteProductByID(productID int64)bool {
	return p.productRepository.Delete(productID)
}

// 插入函数
func (p *ProductService) InsertProduct(product *datamodels.Product) (int64, error) {
	return p.productRepository.Insert(product)
}

// 更新函数
func (p *ProductService) UpdateProduct(product *datamodels.Product)error  {
	return p.productRepository.Update(product)
}

