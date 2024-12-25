package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/service"
	"github.com/onlylight29/go-ecommerce-backend-api/pkg/response"
)

// type UserController struct {
// 	userService *service.UserService
// }

// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }

// func (uc *UserController) GetUserByID(c *gin.Context) {
// 	// name := c.DefaultQuery("name", "unknown")
// 	// uid := c.Param("uid")

// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	"message": uc.userService.GetInfoUserService(),
// 	// 	"name":    name,
// 	// 	"uid":     uid,
// 	// })

// 	response.SuccessResponse(c, response.ErrCodeSuccess, uc.userService.GetInfoUserService())
// }

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	email := c.DefaultQuery("email", "unknown")
	purpose := c.DefaultQuery("purpose", "unknown")

	response.SuccessResponse(c, uc.userService.Register(email, purpose), nil)
}
