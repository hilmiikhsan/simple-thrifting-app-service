package product

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/thrifting-app-service/constants"
	"github.com/hilmiikhsan/thrifting-app-service/helpers"
	"github.com/hilmiikhsan/thrifting-app-service/internal/dto"
	"github.com/hilmiikhsan/thrifting-app-service/internal/interfaces"
	"github.com/hilmiikhsan/thrifting-app-service/internal/validator"
)

type ProductHandler struct {
	ProductService interfaces.IProductService
	Validator      *validator.Validator
}

func (api *ProductHandler) CreateProduct(ctx *gin.Context) {
	var (
		req = new(dto.CreateProductRequest)
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.Logger.Error("handler::Login - Failed to bind request : ", err)
		ctx.JSON(http.StatusBadRequest, helpers.Error(constants.ErrFailedBadRequest))
		return
	}

	if err := api.Validator.Validate(req); err != nil {
		helpers.Logger.Error("handler::Login - Failed to validate request : ", err)
		code, errs := helpers.Errors(err, req)
		ctx.JSON(code, helpers.Error(errs))
		return
	}

	err := api.ProductService.CreateProduct(ctx.Request.Context(), req)
	if err != nil {
		helpers.Logger.Error("handler::CreateProduct - Failed to create product : ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(nil, ""))
}

func (api *ProductHandler) GetAllProduct(ctx *gin.Context) {
	products, err := api.ProductService.GetAllProduct(ctx.Request.Context())
	if err != nil {
		helpers.Logger.Error("handler::GetAllProduct - Failed to get all product : ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(products, ""))
}

func (api *ProductHandler) GetDetailProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := api.ProductService.GetDetailProduct(ctx.Request.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrProductNotFound) {
			ctx.JSON(http.StatusNotFound, helpers.Error(err))
			return
		}

		helpers.Logger.Error("handler::GetDetailProduct - Failed to get product by id : ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.Error(err))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(product, ""))
}
