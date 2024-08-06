package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"github.com/dilshodforever/fooddalivary-api/api/handler"
	_ "github.com/dilshodforever/fooddalivary-api/docs"

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Gateway
// @version 1.0
// @description API documentation for the gateway service
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	// Middleware setup if needed
	ca, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		panic(err)
	}

	err = ca.LoadPolicy()
	if err != nil {
		panic(err)
	}

	router := r.Group("/")

	// Swagger documentation
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))

	// Product endpoints
	p := router.Group("/products")
	{
		p.POST("/", h.CreateProduct)
		p.GET("/:id", h.GetProduct)
		p.PUT("/", h.UpdateProduct)
		p.DELETE("/:id", h.DeleteProduct)
		p.GET("/", h.ListProducts)
	}

	// OrderItem endpoints
	o := router.Group("/orderitems")
	{
		o.POST("/", h.AddItems)
		o.DELETE("/", h.DeleteItems)
		o.GET("/", h.ListOrderItems)
	}
	or := router.Group("/orders")
	{
		or.POST("/", h.CreateOrder)
		or.GET("/:id", h.GetOrderById)
		or.PUT("/", h.UpdateOrder)
		or.DELETE("/:id", h.DeleteOrder)
		or.GET("/", h.ListOrders)
		or.PATCH("/status", h.UpdateStatus)
	}

	return r
}
