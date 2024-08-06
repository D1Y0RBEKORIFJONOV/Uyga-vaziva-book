package router

import (
	handler2 "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/api/handler"
	bookusecase "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/usecase/book"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRouter(impl *bookusecase.BookUseCaseImpl) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	handler := handler2.NewBookHandler(impl)
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	bookGroup := r.Group("/book")
	{
		bookGroup.POST("/create", handler.CreateBook)
		bookGroup.PUT("/update", handler.UpdateBook)
		bookGroup.DELETE("/delete", handler.DeleteBook)
		bookGroup.GET("/list", handler.GetBooks)
		bookGroup.GET("", handler.GetBook)
	}
	return r
}
