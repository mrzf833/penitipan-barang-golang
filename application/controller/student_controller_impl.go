package controller

import (
	"net/http"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/model/web"
	"penitipan-barang/application/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type StudentControllerImpl struct {
	StudentService service.StudentService
}

func NewStudentController(studentService service.StudentService) StudentController {
	return &StudentControllerImpl{
		StudentService: studentService,
	}
}

func (controller *StudentControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	studentCreateRequest := web.StudentCreateRequest{}

	helper.ReadFromRequestBody(r, &studentCreateRequest)

	studentResponse := controller.StudentService.Create(r.Context(), studentCreateRequest)
	webResponse := web.WebResponse{
		Message: "Success create student",
		Data:    studentResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *StudentControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	studentId, err := strconv.Atoi(params.ByName("studentId"))
	helper.PanicIfError(err)

	studentUpdateRequest := web.StudentUpdateRequest{}

	helper.ReadFromRequestBody(r, &studentUpdateRequest)

	studentUpdateRequest.Id = studentId

	studentResponse := controller.StudentService.Update(r.Context(), studentUpdateRequest)
	webResponse := web.WebResponse{
		Message: "Success update student",
		Data:    studentResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *StudentControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	studentId, err := strconv.Atoi(params.ByName("studentId"))
	helper.PanicIfError(err)

	controller.StudentService.Delete(r.Context(), studentId)
	webResponse := web.WebResponse{
		Message: "Success delete student",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *StudentControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	studentId, err := strconv.Atoi(params.ByName("studentId"))
	helper.PanicIfError(err)

	student := controller.StudentService.FindById(r.Context(), studentId)
	webResponse := web.WebResponse{
		Message: "Success get student",
		Data:    student,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *StudentControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	students := controller.StudentService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Message: "Success get all student",
		Data:    students,
	}

	helper.WriteToResponseBody(w, webResponse)
}
