package controller

import (
	"net/http"
	"penitipan-barang/application/helper"
	"penitipan-barang/application/model/web"
	"penitipan-barang/application/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type InventoryControllerImpl struct {
	InventoryService service.InventoryService
}

func NewInventoryController(inventoryService service.InventoryService) InventoryController {
	return &InventoryControllerImpl{
		InventoryService: inventoryService,
	}
}

func (controller *InventoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	inventoryCreateRequest := web.InventoryCreateRequest{}

	helper.ReadFromRequestBody(r, &inventoryCreateRequest)

	inventoryResponse := controller.InventoryService.Create(r.Context(), inventoryCreateRequest)
	webResponse := web.WebResponse{
		Message: "Success create inventory",
		Data:    inventoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *InventoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	inventoryId, err := strconv.Atoi(params.ByName("inventoryId"))
	helper.PanicIfError(err)

	inventoryUpdateRequest := web.InventoryUpdateRequest{}

	helper.ReadFromRequestBody(r, &inventoryUpdateRequest)

	inventoryUpdateRequest.Id = inventoryId

	inventoryResponse := controller.InventoryService.Update(r.Context(), inventoryUpdateRequest)
	webResponse := web.WebResponse{
		Message: "Success update inventory",
		Data:    inventoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *InventoryControllerImpl) TakeInventory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	inventoryId, err := strconv.Atoi(params.ByName("inventoryId"))
	helper.PanicIfError(err)

	inventoryTakeRequest := web.InventoryTakeRequest{}

	helper.ReadFromRequestBody(r, &inventoryTakeRequest)

	inventoryTakeRequest.Id = inventoryId

	inventoryResponse := controller.InventoryService.TakeInventory(r.Context(), inventoryTakeRequest)
	webResponse := web.WebResponse{
		Message: "Success take inventory",
		Data:    inventoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *InventoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	inventoryId, err := strconv.Atoi(params.ByName("inventoryId"))
	helper.PanicIfError(err)

	controller.InventoryService.Delete(r.Context(), inventoryId)
	webResponse := web.WebResponse{
		Message: "Success delete inventory",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *InventoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	inventoryId, err := strconv.Atoi(params.ByName("inventoryId"))
	helper.PanicIfError(err)

	inventory := controller.InventoryService.FindById(r.Context(), inventoryId)
	webResponse := web.WebResponse{
		Message: "Success get inventory",
		Data:    inventory,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *InventoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	inventories := controller.InventoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Message: "Success get all inventory",
		Data:    inventories,
	}

	helper.WriteToResponseBody(w, webResponse)
}
