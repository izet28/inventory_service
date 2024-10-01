package controllers

import (
	"encoding/json"
	"inventory_service/models"
	"inventory_service/services"
	"inventory_service/utils"
	"net/http"

	"github.com/sirupsen/logrus"
)

type WarehouseController struct {
	service services.WarehouseService
	logger  *logrus.Logger
}

// NewProductController membuat instance baru dari ProductController
func NewProductController(service services.WarehouseService, logger *logrus.Logger) *WarehouseController {
	return &WarehouseController{
		service: service,
		logger:  logger,
	}
}
func (c *WarehouseController) CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	var warehouse models.Warehouse
	if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
		appErr := utils.NewAppError(http.StatusBadRequest, "Invalid input", "Input yang anda masukan salah", err)
		c.logger.Warn("input tidak valid: ", err)
		utils.RespondWithError(w, appErr)
		return
	}

}
