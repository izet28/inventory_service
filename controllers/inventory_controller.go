package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "inventory_service/models"
    "inventory_service/services"
    "inventory_service/utils"
    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"
)

// InventoryController menangani permintaan HTTP untuk inventory
type InventoryController struct {
    service services.InventoryService
    logger  *logrus.Logger
}

// NewInventoryController membuat instance baru dari InventoryController
func NewInventoryController(service services.InventoryService, logger *logrus.Logger) *InventoryController {
    return &InventoryController{
        service: service,
        logger:  logger,
    }
}

// GetInventory mengembalikan stok berdasarkan ID produk
func (c *InventoryController) GetInventory(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    productID, err := strconv.Atoi(params["product_id"])
    if err != nil {
        appErr := utils.NewAppError(http.StatusBadRequest, "INVALID_PRODUCT_ID", "ID produk tidak valid", err)
        c.logger.Warn(appErr)
        utils.RespondWithError(w, appErr)
        return
    }

    inventory, appErr := c.service.GetInventory(uint(productID))
    if appErr != nil {
        c.logger.Error(appErr)
        utils.RespondWithError(w, appErr)
        return
    }

    utils.RespondJSON(w, http.StatusOK, inventory)
}

// UpdateStock memperbarui stok produk
func (c *InventoryController) UpdateStock(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    productID, err := strconv.Atoi(params["product_id"])
    if err != nil {
        appErr := utils.NewAppError(http.StatusBadRequest, "INVALID_PRODUCT_ID", "ID produk tidak valid", err)
        c.logger.Warn(appErr)
        utils.RespondWithError(w, appErr)
        return
    }

    var reqData struct {
        NewStock int json:"new_stock"
    }
    if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
        appErr := utils.NewAppError(http.StatusBadRequest, "INVALID_REQUEST_DATA", "Data request tidak valid", err)
        c.logger.Warn(appErr)
        utils.RespondWithError(w, appErr)
        return
    }

    appErr := c.service.UpdateStock(uint(productID), reqData.NewStock)
    if appErr != nil {
        c.logger.Error(appErr)
        utils.RespondWithError(w, appErr)
        return
    }

    utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Stok diperbarui"})
}

// ReplenishStock menambah stok produk
func (c *InventoryController) ReplenishStock(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    productID, err := strconv.Atoi(params["product_id"])
    if err != nil {
        appErr := utils.NewAppError(http.StatusBadRequest, "INVALID_PRODUCT_ID", "ID produk tidak valid", err)
        c.logger.Warn(appErr)
        utils.RespondWithError(w, appErr)
        return
    }

    var reqData struct {
        AdditionalStock int json:"additional_stock"
    }
    if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
        appErr := utils.NewAppError(http.StatusBadRequest, "INVALID_REQUEST_DATA", "Data request tidak valid", err)
        c.logger.Warn(appErr)
        utils.RespondWithError(w, appErr)
        return
    }

    appErr := c.service.ReplenishStock(uint(productID), reqData.AdditionalStock)
    if appErr != nil {
        c.logger.Error(appErr)
        utils.RespondWithError(w, appErr)
        return
    }

    utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Stok ditambah"})
}