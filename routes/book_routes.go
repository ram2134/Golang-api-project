package routes

import (
    "github.com/gin-gonic/gin"
    "bookstore-api/controller"
    "bookstore-api/middleware"
)

func RegisterBookRoutes(r *gin.Engine, ctrl *controller.BookController) {
    r.POST("/books", middleware.ErrorHandler(ctrl.CreateBook))
    r.GET("/books", middleware.ErrorHandler(ctrl.GetBooks))
    r.GET("/books/:id", middleware.ErrorHandler(ctrl.GetBook))
    r.PUT("/books/:id", middleware.ErrorHandler(ctrl.UpdateBook))
    r.DELETE("/books/:id", middleware.ErrorHandler(ctrl.DeleteBook))
}
