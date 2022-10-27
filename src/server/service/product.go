package service

import (
	"database/sql"
	"mini_e_commerce/src/server/params"
	"mini_e_commerce/src/server/repository"
	"mini_e_commerce/src/server/view"
)

type ProductService struct {
	repo         repository.ProductRepo
	categoryRepo repository.CategoryRepo
}

func NewProductServices(repo repository.ProductRepo, categoryRepo repository.CategoryRepo) *ProductService {
	return &ProductService{
		repo:         repo,
		categoryRepo: categoryRepo,
	}
}

func (p *ProductService) GetProducts() *view.Response {
	products, err := p.repo.GetProducts()
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFindAll(view.NewProductGetAllResponse(products))
}

func (p *ProductService) CreateProduct(req *params.ProductCreate) *view.Response {
	product := req.ParseToModel()
	category, err := p.categoryRepo.GetCategoryById(product.CategoryId)
	if err != nil {
		return view.ErrBadRequest("category id doesn't exists")
	}

	product.Category = *category
	err = p.repo.CreateProduct(product)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessCreated(view.NewProductCreateResponse(product))
}

func (p *ProductService) UpdateProduct(productId string, req *params.ProductUpdate) *view.Response {
	product := req.ParseToModel()
	category, err := p.categoryRepo.GetCategoryById(product.CategoryId)
	if err != nil {
		return view.ErrBadRequest("category id doesn't exists")
	}

	getProduct, err := p.repo.GetProductById(productId)
	if err != nil {
		return view.ErrBadRequest("product doesn't exists")
	}

	product.Category = *category
	product.ID = getProduct.ID
	err = p.repo.UpdateProduct(productId, product)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessUpdated(view.NewProductUpdateResponse(product))
}

func (p *ProductService) DeleteProduct(productId string) *view.Response {
	product, err := p.repo.GetProductById(productId)
	if err != nil {
		return view.ErrBadRequest("product doesn't exists")
	}

	err = p.repo.DeleteProduct(productId, product)
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessDeleted(view.NewProductUpdateResponse(product))
}

func (p *ProductService) GetProductById(productId string) *view.Response {
	product, err := p.repo.GetProductById(productId)
	if err != nil {
		return view.ErrInternalServer(err.Error())

	}

	return view.SuccessFindAll(product)
}
