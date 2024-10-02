package controllers

import (
	"encoding/json"
	"inventory_service/models"
	"inventory_service/services"
	"inventory_service/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// WarehouseController menangani permintaan HTTP terkait warehouse
type WarehouseController struct {
	service services.WarehouseService
	logger  *logrus.Logger
}

// NewWarehouseController membuat instance baru dari WarehouseController
func NewWarehouseController(service services.WarehouseService, logger *logrus.Logger) *WarehouseController {
	return &WarehouseController{
		service: service,
		logger:  logger,
	}
}

// CreateWarehouse menangani permintaan untuk membuat warehouse baru
func (c *WarehouseController) CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	var warehouse models.Warehouse
	if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusBadRequest, "INVALID_REQUEST", "Data tidak valid", err))
		return
	}

	c.logger.Infof("Permintaan POST untuk membuat warehouse: %+v", warehouse)

	createdWarehouse, err := c.service.CreateWarehouse(&warehouse)
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusInternalServerError, "SERVER_ERROR", err.Error(), err))
		return
	}

	utils.RespondJSON(w, http.StatusCreated, createdWarehouse)
}

// GetAllWarehouses menangani permintaan untuk mendapatkan semua warehouse
func (c *WarehouseController) GetAllWarehouses(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Permintaan GET untuk semua warehouse diterima")

	warehouses, err := c.service.GetAllWarehouses()
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusInternalServerError, "SERVER_ERROR", "Gagal mengambil semua warehouse", err))
		return
	}

	utils.RespondJSON(w, http.StatusOK, warehouses)
}

// GetWarehouseByID menangani permintaan untuk mendapatkan warehouse berdasarkan ID
func (c *WarehouseController) GetWarehouseByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusBadRequest, "INVALID_ID", "ID warehouse tidak valid", err))
		return
	}

	c.logger.Infof("Permintaan GET untuk warehouse dengan ID: %d", id)

	warehouse, err := c.service.GetWarehouseByID(uint(id))
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusNotFound, "NOT_FOUND", "Warehouse tidak ditemukan", err))
		return
	}

	utils.RespondJSON(w, http.StatusOK, warehouse)
}

// UpdateWarehouse menangani permintaan untuk memperbarui warehouse
func (c *WarehouseController) UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusBadRequest, "INVALID_ID", "ID warehouse tidak valid", err))
		return
	}

	var warehouse models.Warehouse
	if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusBadRequest, "INVALID_REQUEST", "Data tidak valid", err))
		return
	}

	warehouse.ID = uint(id) // Set ID yang ada

	c.logger.Infof("Permintaan PUT untuk warehouse dengan ID: %d", warehouse.ID)

	updatedWarehouse, err := c.service.UpdateWarehouse(&warehouse)
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusInternalServerError, "SERVER_ERROR", err.Error(), err))
		return
	}

	utils.RespondJSON(w, http.StatusOK, updatedWarehouse)
}

// DeleteWarehouse menangani permintaan untuk menghapus warehouse
func (c *WarehouseController) DeleteWarehouse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusBadRequest, "INVALID_ID", "ID warehouse tidak valid", err))
		return
	}

	c.logger.Infof("Permintaan DELETE untuk warehouse dengan ID: %d", id)

	err = c.service.DeleteWarehouse(uint(id))
	if err != nil {
		utils.RespondWithError(w, utils.NewAppError(http.StatusInternalServerError, "SERVER_ERROR", err.Error(), err))
		return
	}

	utils.RespondJSON(w, http.StatusNoContent, nil)
}
