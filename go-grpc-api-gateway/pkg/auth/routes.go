package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/moromin/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/moromin/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	route := r.Group("/auth")
	route.POST("/register", svc.Register)
	route.POST("/login", svc.Login)

	return svc
}

// wrap the methods to use gin
func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
