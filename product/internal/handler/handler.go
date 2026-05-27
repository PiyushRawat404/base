package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"product/internal/models"
	"product/internal/service"
	"product/pkg/utils"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var req models.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	if err := validateRequest(req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	response, err := h.service.Signup(r.Context(), req)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, response)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	if err := validateRequest(req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	response, err := h.service.Login(r.Context(), req)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, response)
}

func (h *Handler) Products(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createProduct(w, r)
	case http.MethodGet:
		h.listProducts(w, r)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (h *Handler) ProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid product id"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		product, err := h.service.GetProductByID(r.Context(), id)
		if err != nil {
			writeServiceError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, product)
	case http.MethodPut:
		var req models.Product
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
			return
		}
		if err := validateRequest(req); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		product, err := h.service.UpdateProduct(r.Context(), id, req)
		if err != nil {
			writeServiceError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, product)
	case http.MethodDelete:
		if err := h.service.DeleteProduct(r.Context(), id); err != nil {
			writeServiceError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func (h *Handler) createProduct(w http.ResponseWriter, r *http.Request) {
	var req models.Product
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}
	if err := validateRequest(req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	product, err := h.service.CreateProduct(r.Context(), req)
	if err != nil {
		writeServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, product)
}

func (h *Handler) listProducts(w http.ResponseWriter, r *http.Request) {
	pagination := utils.ParsePagination(r.URL.Query())
	filter := utils.ParseProductFilter(r.URL.Query())

	response, err := h.service.GetProducts(r.Context(), pagination, filter)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to fetch products"})
		return
	}

	writeJSON(w, http.StatusOK, response)
}

func writeServiceError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, service.ErrValidation),
		errors.Is(err, service.ErrWeakPassword):
		status = http.StatusBadRequest
	case errors.Is(err, service.ErrDuplicateEmail):
		status = http.StatusConflict
	case errors.Is(err, service.ErrInvalidCredential):
		status = http.StatusUnauthorized
	case errors.Is(err, service.ErrProductNotFound):
		status = http.StatusNotFound
	}

	writeJSON(w, status, map[string]string{"error": err.Error()})
}

func validateRequest(payload any) error {
	if err := models.Validate.Struct(payload); err != nil {
		var validationErrs validator.ValidationErrors
		if errors.As(err, &validationErrs) {
			for _, validationErr := range validationErrs {
				switch validationErr.Tag() {
				case "required":
					return errors.New(validationErr.Field() + " is required")
				case "email":
					return errors.New("email must be a valid email address")
				case "strong_password":
					return service.ErrWeakPassword
				case "min":
					return errors.New(validationErr.Field() + " is too short")
				case "max":
					return errors.New(validationErr.Field() + " is too long")
				case "gt":
					return errors.New(validationErr.Field() + " must be greater than 0")
				case "gte":
					return errors.New(validationErr.Field() + " must be 0 or greater")
				}
			}
		}
		return service.ErrValidation
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
