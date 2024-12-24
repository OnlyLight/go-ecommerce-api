package routers

import (
	"github.com/onlylight29/go-ecommerce-backend-api/internal/routers/manage"
	"github.com/onlylight29/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = &RouterGroup{}
