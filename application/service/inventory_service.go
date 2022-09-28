package service

import (
	"context"
	"penitipan-barang/application/model/web"
)

type InventoryService interface {
	Create(ctx context.Context, request web.InventoryCreateRequest) web.InventoryResponse
	Update(ctx context.Context, request web.InventoryUpdateRequest) web.InventoryResponse
	TakeInventory(ctx context.Context, request web.InventoryTakeRequest) web.InventoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, inventoryId int) web.InventoryResponse
	FindAll(ctx context.Context) []web.InventoryResponse
}
