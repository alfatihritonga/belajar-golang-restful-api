package app

import (
	"github.com/alfatihritonga/belajar-golang-restful-api/contoller"
	"github.com/alfatihritonga/belajar-golang-restful-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController contoller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
