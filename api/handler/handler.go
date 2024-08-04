package handler

import (
	_ "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/api/docs"
	bookeentity "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/entity/book"
	bookusecase "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/usecase/book"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookHandler struct {
	book bookusecase.BookUseCase
}

func NewBookHandler(book bookusecase.BookUseCase) *BookHandler {
	return &BookHandler{
		book: book,
	}
}

// CreateBook godoc
// @Summary createBook
// @Description New create book
// @Tags book
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body bookEntity.CreateBookRequest true "Create new user book"
// @Success 200 {object} bookEntity.Book
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /book/create [post]
func (b *BookHandler) CreateBook(c *gin.Context) {
	var req bookeentity.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book, err := b.book.CreateBook(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// UpdateBook godoc
// @Summary createBook
// @Description New create book
// @Tags book
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body bookEntity.UpdateBookRequest true "Create new user book"
// @Success 200 {object} bookEntity.Book
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /book/update [put]
func (b *BookHandler) UpdateBook(c *gin.Context) {
	var req bookeentity.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book, err := b.book.UpdateBook(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// DeleteBook godoc
// @Summary createBook
// @Description New create book
// @Tags book
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param field query string true "field"
// @Param value query string true "value"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /book/delete [delete]
func (b *BookHandler) DeleteBook(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	if field == "" || value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field or value is empty"})
		return
	}
	err := b.book.DeleteBook(c.Request.Context(), &bookeentity.FieldValueReq{
		Field: field,
		Value: value,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{
		"status": "deleted",
	})
}

// GetBook godoc
// @Summary GetBook
// @Description New create book
// @Tags book
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param field query string true "field"
// @Param value query string true "value"
// @Success 200 {object} bookEntity.Book
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /book [get]
func (b *BookHandler) GetBook(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	if field == "" || value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field or value is empty"})
		return
	}
	book, err := b.book.GetBook(c.Request.Context(), &bookeentity.FieldValueReq{
		Field: field,
		Value: value,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

// GetBooks godoc
// @Summary GetBooks
// @Description New create book
// @Tags book
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param field query string true "field"
// @Param value query string true "value"
// @Param page query int true "Page number"
// @Param limit query int true "Number of items per page"
// @Success 200 {object} []bookEntity.Book
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /book/ [get]
func (b *BookHandler) GetBooks(c *gin.Context) {
	field := c.Query("field")
	value := c.Query("value")
	page := c.Query("page")
	limit := c.Query("limit")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books, err := b.book.GetBookAll(c.Request.Context(), &bookeentity.GetAllBookReq{
		Field:  field,
		Value:  value,
		Limit:  limitInt,
		Offset: pageInt,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
