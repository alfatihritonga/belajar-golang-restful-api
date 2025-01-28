package service

import (
	"context"
	"database/sql"

	"github.com/alfatihritonga/belajar-golang-restful-api/exception"
	"github.com/alfatihritonga/belajar-golang-restful-api/helper"
	"github.com/alfatihritonga/belajar-golang-restful-api/model/domain"
	"github.com/alfatihritonga/belajar-golang-restful-api/model/web"
	"github.com/alfatihritonga/belajar-golang-restful-api/repository"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (c *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := c.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = c.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := c.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := c.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = c.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := c.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	c.CategoryRepository.Delete(ctx, tx, category)
}

func (c *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := c.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := c.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
