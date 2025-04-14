package controller

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"projeto-magalu-api-go/pkg/api/model"
	"projeto-magalu-api-go/pkg/api/service"
)

type ProductCtrl interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
}

type ProductCtrlImpl struct {
	service service.Product
}

func NewProductController(service service.Product) *ProductCtrlImpl {
	return &ProductCtrlImpl{service: service}
}

// CreateProduct handles the creation of a new product
// swagger:route POST /products createProduct
// Create a new product with the input payload
// responses:
//   201: ProductOut
//   400: ErrorResponse
//   422: ErrorResponse
// responses:
//
//	201: ProductResponse
func (p *ProductCtrlImpl) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var in model.ProductIn
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	error := ValidateProduct(&in)
	if error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	product, err := p.service.Create(context.Background(), &in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(w)
	encoder.Encode(product)
}

// GetProduct handles fetching the details of a product by ID
// swagger:route GET /products/{productId} products getProduct
// Get product details by product ID
// responses:
//   200: ProductOut
//   404: ErrorResponse
//   400: ErrorResponse
// responses:
//
//	200: ProductResponse
//	404: ErrorResponse
func (p *ProductCtrlImpl) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	product, err := p.service.GetByID(context.Background(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func ValidateProduct(v interface{}) error {
	var validate *validator.Validate
	validate = validator.New()

	errs := validate.Struct(v)
	if errs != nil {
		return errs
	}
	return nil
}
