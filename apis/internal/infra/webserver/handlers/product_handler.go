package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/waldrey/go-expert/apis/internal/dto"
	"github.com/waldrey/go-expert/apis/internal/entity"
	"github.com/waldrey/go-expert/apis/internal/infra/database"
	entityPkg "github.com/waldrey/go-expert/apis/pkg/entity"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// CreateProduct godoc
// @Summary 	Create a product
// @Description Create a product
// @Tags 		products
// @Accept  	json
// @Produce  	json
// @Param 		request		product 	body 		dto.CreateProductInput	true	"product request"
// @Success 	201 	{object} 	string
// @Failure 	500 	{object} 	Error
// @Router 		/products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetProduct godoc
// @Summary 	Get a product
// @Description Get a product
// @Tags 		products
// @Accept  	json
// @Produce  	json
// @Param 		id		path 	string 	true	"Product ID" Format(uuid)
// @Success 	200 	{object} 	entity.Product
// @Failure 	404
// @Failure 	500 	{object} 	Error
// @Router 		/products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// GetProducts godoc
// @Summary 	Get a products
// @Description Get a products
// @Tags 		products
// @Accept  	json
// @Produce  	json
// @Param 		page	path 	string 	false	"page number"
// @Param 		limit	path 	string 	false	"limit"
// @Success 	200 	{object} 	entity.Product
// @Failure 	404 	{object} 	Error
// @Failure 	500 	{object} 	Error
// @Router 		/products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	limit := chi.URLParam(r, "limit")
	sort := chi.URLParam(r, "sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// UpdateProduct godoc
// @Summary 	Update a product
// @Description Update a product
// @Tags 		products
// @Accept  	json
// @Produce  	json
// @Param 		id		path 	string 	true	"Product ID" Format(uuid)
// @Param 		request		product 	body 		dto.CreateProductInput	true	"product request"
// @Success 	200
// @Failure 	404
// @Failure 	500 	{object} 	Error
// @Router 		/products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteProduct godoc
// @Summary 	Delete a product
// @Description Delete a product
// @Tags 		products
// @Accept  	json
// @Produce  	json
// @Param 		id		path 	string 	true	"Product ID" Format(uuid)
// @Success 	200
// @Failure 	404
// @Failure 	500 	{object} 	Error
// @Router 		/products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
