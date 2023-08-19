package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeypc/go-jwt-mux/helper"
	"github.com/jeypc/go-jwt-mux/models"
)

func GetDataAll(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	if err := models.DB.Find(&products).Error; err != nil {
		res := map[string]string{"status": "500", "message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, res)
		return
	}

	helper.ResponseJSON(w, http.StatusOK, products)
}

func FindOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productIDStr := params["id"]

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		res := map[string]string{"status": "400", "message": "Invalid product ID"}
		helper.ResponseJSON(w, http.StatusBadRequest, res)
		return
	}

	var product models.Product
	if err := models.DB.First(&product, productID).Error; err != nil {
		res := map[string]string{"status": "404", "message": "Product not found"}
		helper.ResponseJSON(w, http.StatusNotFound, res)
		return
	}

	helper.ResponseJSON(w, http.StatusOK, product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newProduct); err != nil {
		res := map[string]string{"status": "400", "message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, res)
		return
	}
	defer r.Body.Close()

	if err := models.DB.Create(&newProduct).Error; err != nil {
		res := map[string]string{"status": "500", "message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, res)
		return
	}

	res := map[string]string{"status": "201", "message": "Product created successfully"}
	helper.ResponseJSON(w, http.StatusCreated, res)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productIDStr := params["id"]

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		res := map[string]string{"status": "400", "message": "Invalid product ID"}
		helper.ResponseJSON(w, http.StatusBadRequest, res)
		return
	}

	var updateProduct models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updateProduct); err != nil {
		res := map[string]string{"status": "400", "message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, res)
		return
	}
	defer r.Body.Close()

	updateProduct.Id = int64(productID)
	if err := models.DB.Save(&updateProduct).Error; err != nil {
		res := map[string]string{"status": "500", "message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, res)
		return
	}

	res := map[string]string{"status": "200", "message": "Product updated successfully"}
	helper.ResponseJSON(w, http.StatusOK, res)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productIDStr := params["id"]

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		res := map[string]string{"status": "400", "message": "Invalid product ID"}
		helper.ResponseJSON(w, http.StatusBadRequest, res)
		return
	}

	if err := models.DB.Delete(&models.Product{}, productID).Error; err != nil {
		res := map[string]string{"status": "500", "message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, res)
		return
	}

	res := map[string]string{"status": "200", "message": "Product deleted successfully"}
	helper.ResponseJSON(w, http.StatusOK, res)
}
