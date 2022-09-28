package controller

import (
	"net/http"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/model/web"
	"penitipan-barang/application/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}

	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Message: "Success create category",
		Data:    categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	categoryUpdateRequest := web.CategoryUpdateRequest{}

	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	categoryUpdateRequest.Id = categoryId

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Message: "Success update category",
		Data:    categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), categoryId)
	webResponse := web.WebResponse{
		Message: "Success delete category",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	category := controller.CategoryService.FindById(r.Context(), categoryId)
	webResponse := web.WebResponse{
		Message: "Success get category",
		Data:    category,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categories := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Message: "Success get all category",
		Data:    categories,
	}

	helper.WriteToResponseBody(w, webResponse)
}
