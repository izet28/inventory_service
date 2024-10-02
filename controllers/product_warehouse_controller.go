package controllers

import (
	"encoding/json"
	"inventory_service/services"
	"inventory_service/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ProductWarehouseController menangani stok produk di warehouse
type ProductWarehouseController struct {
	service services.ProductWarehouseService
}

// NewProductWarehouseController membuat instance baru dari ProductWarehouseController
func NewProductWarehouseController(service services.ProductWarehouseService) *ProductWarehouseController {
	return &ProductWarehouseController{service: service}
}

// AddStockToWarehouse menambahkan stok produk ke warehouse
func (c *ProductWarehouseController) AddStockToWarehouse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID, err := strconv.Atoi(params["product_id"])
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusBadRequest, "INVALID_PRODUCT_ID", "ID produk tidak valid", err))
		return
	}
	warehouseID, err := strconv.Atoi(params["warehouse_id"])
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusBadRequest, "INVALID_WAREHOUSE_ID", "ID warehouse tidak valid", err))
		return
	}

	var requestBody struct {
		Stock int `json:"stock"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusBadRequest, "INVALID_REQUEST", "Data tidak valid", err))
		return
	}

	err = c.service.AddStockToWarehouse(uint(productID), uint(warehouseID), requestBody.Stock)
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusInternalServerError, "SERVER_ERROR", "Gagal menambahkan stok", err))
		return
	}

	utils.RespondJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Stok berhasil ditambahkan",
	})
}

// UpdateStock memperbarui stok produk di warehouse
func (c *ProductWarehouseController) UpdateStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID, err := strconv.Atoi(params["product_id"])
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusBadRequest, "INVALID_PRODUCT_ID", "ID produk tidak valid", err))
		return
	}
	warehouseID, err := strconv.Atoi(params["warehouse_id"])
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusBadRequest, "INVALID_WAREHOUSE_ID", "ID warehouse tidak valid", err))
		return
	}

	var requestBody struct {
		Stock int `json:"stock"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusBadRequest, "INVALID_REQUEST", "Data tidak valid", err))
		return
	}

	err = c.service.UpdateStock(uint(productID), uint(warehouseID), requestBody.Stock)
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusInternalServerError, "SERVER_ERROR", "Gagal memperbarui stok", err))
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Stok berhasil diperbarui",
	})
}

// GetStock mendapatkan stok produk di warehouse tertentu
func (c *ProductWarehouseController) GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID, err := strconv.Atoi(params["product_id"])
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusBadRequest, "INVALID_PRODUCT_ID", "ID produk tidak valid", err))
		return
	}
	warehouseID, err := strconv.Atoi(params["warehouse_id"])
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusBadRequest, "INVALID_WAREHOUSE_ID", "ID warehouse tidak valid", err))
		return
	}

	stock, err := c.service.GetStockByProductAndWarehouse(uint(productID), uint(warehouseID))
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusNotFound, "NOT_FOUND", "Stok tidak ditemukan", err))
		return
	}

	utils.RespondJSON(w, http.StatusOK, stock)
}
