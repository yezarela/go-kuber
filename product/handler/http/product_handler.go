package http

import (
	"net/http"
	"strconv"

	"kuba/models"
	_productUsecase "kuba/product/usecase"
	"kuba/utils/tern"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	usecase _productUsecase.Usecase
}

func NewProductHandler(e *echo.Echo, uc _productUsecase.Usecase) {
	handler := &ProductHandler{
		usecase: uc,
	}

	e.GET("/products/:id", handler.GetProductByID)
	e.GET("/products", handler.ListProduct)
	e.POST("/products", handler.CreateProduct)
}

func (m *ProductHandler) GetProductByID(c echo.Context) error {
	paramId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Not found")
	}

	id := uint(paramId)
	ctx := c.Request().Context()

	product, err := m.usecase.GetProductByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

func (m *ProductHandler) ListProduct(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	limit = tern.Int(limit, 10)

	ctx := c.Request().Context()

	products, err := m.usecase.ListProduct(ctx, limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, products)
}

type CreateProductBody struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (m *ProductHandler) CreateProduct(c echo.Context) error {
	body := CreateProductBody{}

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()

	data := models.Product{
		Name:        body.Name,
		Description: body.Description,
	}

	product, err := m.usecase.CreateProduct(ctx, data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}
