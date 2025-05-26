package controller

import (
	"net/http"

	"FP-DevOps/config"
	"FP-DevOps/constants"
	"FP-DevOps/dto"
	"FP-DevOps/entity"
	"FP-DevOps/service"
	"FP-DevOps/utils"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		Me(ctx *gin.Context)
	}

	userController struct {
		jwtService  config.JWTService
		userService service.UserService
	}
)

func NewUserController(us service.UserService, jwt config.JWTService) UserController {
	return &userController{
		jwtService:  jwt,
		userService: us,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var user dto.UserRequest
	if err := ctx.ShouldBind(&user); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.userService.RegisterUser(ctx.Request.Context(), user)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REGISTER_USER, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userController) Login(ctx *gin.Context) {
	var req dto.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.userService.Login(ctx.Request.Context(), req.Username, req.Password)
	if err != nil {
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	token := c.jwtService.GenerateToken(user.ID.String(), user.Username)
	userResponse := entity.Authorization{
		Token:    token,
		Username: user.Username,
	}

	response := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_LOGIN, userResponse)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) Me(ctx *gin.Context) {
	userID := ctx.MustGet(constants.CTX_KEY_USER_ID).(string)

	result, err := c.userService.Me(ctx.Request.Context(), userID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_USER, result)
	ctx.JSON(http.StatusOK, res)
}
