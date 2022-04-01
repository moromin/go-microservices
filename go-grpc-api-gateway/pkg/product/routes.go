package product

import (
	"github.com/gin-gonic/gin"
	"github.com/moromin/go-grpc-api-gateway/pkg/auth"
	"github.com/moromin/go-grpc-api-gateway/pkg/config"
	"github.com/moromin/go-grpc-api-gateway/pkg/product/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	route := r.Group("/product")
	route.Use(a.AuthRequired)
	route.POST("/", svc.CreateProduct)
	route.GET("/:id", svc.FindOne)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}
