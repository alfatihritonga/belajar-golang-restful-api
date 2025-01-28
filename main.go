package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/alfatihritonga/belajar-golang-restful-api/app"
	"github.com/alfatihritonga/belajar-golang-restful-api/contoller"
	"github.com/alfatihritonga/belajar-golang-restful-api/helper"
	"github.com/alfatihritonga/belajar-golang-restful-api/middleware"
	"github.com/alfatihritonga/belajar-golang-restful-api/repository"
	"github.com/alfatihritonga/belajar-golang-restful-api/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := contoller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
