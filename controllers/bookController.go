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

	ctx.JSON(http.StatusOK, book)
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

// func UpdateBook(ctx *gin.Context){
// 	bookId := ctx.Param("id")
// 	idData,_ := strconv.Atoi(bookId)
// 	state := false
// 	var updatedBook Book

// 		if err := ctx.ShouldBindJSON(&updatedBook); err != nil{
// 				ctx.AbortWithError(http.StatusBadRequest, err)
// 				return
// 			}

	
// 	for i, book := range bookDatas{
// 		if idData == book.BookID{
// 			state = true
// 			bookDatas[i] = updatedBook
// 			bookDatas[i].BookID = idData
// 			break
// 		}
// 	}

// 	if !state{
// 		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
// 			"error_message": fmt.Sprintf("Book with id %v not found", idData),
// 			})
// 	} else {
// 		show := []byte(`"Updated"`)
// 		ctx.Data(http.StatusOK,"application/json", show)
// 	}
	
// }


// func DeleteBook(ctx *gin.Context){
// 	bookId := ctx.Param("id")
// 	idData,_ := strconv.Atoi(bookId)
// 	state := false
// 	var indexBook int


	
// 	for i, book := range bookDatas{
// 		if idData == book.BookID{
// 			state = true
// 			indexBook = i
// 			break
// 		}
// 	}

// 	if !state{
// 		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
// 			"error_message": fmt.Sprintf("Book with id %v not found", idData),
// 			})
// 		return
// 	}



// 	if state {
// 		copy(bookDatas[indexBook:], bookDatas[indexBook+1:])
// 		bookDatas[len(bookDatas)-1] = Book{}
// 		bookDatas = bookDatas[:len(bookDatas)-1]

// 		show := []byte(`"Deleted"`)
// 		ctx.Data(http.StatusOK,"application/json", show)
// 	}

// 	}


