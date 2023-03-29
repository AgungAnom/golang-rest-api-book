package controllers

import (
	"golang-rest-api-book/models"

	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func (c *Controllers) GetAllBook(ctx *gin.Context){
	book := []models.Book{}
	err := c.projectDB.Find(&book).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK,book)
}

func (c *Controllers) CreateBook(ctx *gin.Context){
	var newBook models.Book
	if err := ctx.ShouldBindJSON(&newBook); err != nil{
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := models.Book{
		Title: newBook.Title,
		Author: newBook.Author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := c.projectDB.Create(&book).Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

func (c *Controllers) GetBook(ctx *gin.Context){
	bookId := ctx.Param("id")
	idData,_ := strconv.Atoi(bookId)
	var book models.Book

	err := c.projectDB.First(&book, idData).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Book with id %v not found", idData),
			})
		return
	}
	ctx.JSON(http.StatusOK,book)
	
}

func (c *Controllers) UpdateBook(ctx *gin.Context){
	bookId := ctx.Param("id")
	idData,_ := strconv.Atoi(bookId)
	var newData models.Book
	var book models.Book
	var bookShow models.Book


	
	err := c.projectDB.First(&book, idData).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Book with id %v not found", idData),
			})
		return
	}
	
	
	if err := ctx.ShouldBindJSON(&newData); err != nil{
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book = models.Book{
		ID:uint(idData),
		Title: newData.Title,
		Author: newData.Author,
		UpdatedAt: time.Now(),
	}

	if err := c.projectDB.Model(&book).Updates(book).Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err2 := c.projectDB.First(&bookShow, idData).Error
	if err2 != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Book with id %v not found", idData),
			})
		return
	}

	ctx.JSON(http.StatusOK, bookShow)

}


func (c *Controllers)DeleteBook(ctx *gin.Context){
	bookId := ctx.Param("id")
	idData,_ := strconv.Atoi(bookId)
	var book models.Book

	err := c.projectDB.First(&book, idData).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Book with id %v not found", idData),
			})
		return
	}

	if err := c.projectDB.Delete(&book).Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"message":"Book deleted successfully",
	})
}


