package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rafaeljpc/golang-first-study/internal/domain/model"
	"github.com/rafaeljpc/golang-first-study/internal/domain/service"

	_ "github.com/rafaeljpc/golang-first-study/docs"
)

type Handler struct {
	service *service.Service
}

type ProductResponse struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ListProductResponse struct {
	result []ProductResponse
}

func NewHttpServiceHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(server *echo.Echo) {
	server.GET("/product", h.ListProducts)
}

// ListProducts returns a list of products.
//	@Summary		List all products
//	@Description	List all products in the system.
//	@Tags			products
//	@Produce		json
//	@Success		200	{object}	ListProductResponse
func (h *Handler) ListProducts(echoCtx echo.Context) error {
	products := convertProducts(h.service.ListProducts())

	response := ListProductResponse{
		result: products,
	}

	return echoCtx.JSON(http.StatusOK, response)
}

func convertProducts(products []model.Product) []ProductResponse {
	result := make([]ProductResponse, len(products))
	for i := range products {
		result[i] = ProductResponse{
			ID:    products[i].ID,
			Name:  products[i].Name,
			Price: products[i].Price,
		}
	}
	return result
}
