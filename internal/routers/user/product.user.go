package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productPublicRouter := Router.Group("/product")
	{
		productPublicRouter.GET("/search")
		productPublicRouter.GET("/detail/:id")
	}
}
