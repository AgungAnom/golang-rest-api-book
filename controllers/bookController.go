package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book	struct{
	BookID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Desc string `json:"desc"`

}
var bookDatas = []Book{}

func CreateBook(ctx *gin.Context){
	var newBook Book
	if err := ctx.ShouldBindJSON(&newBook); err != nil{
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = len(bookDatas)+1
	bookDatas = append(bookDatas,newBook)


	show := []byte(`"Created"`)
	ctx.Data(http.StatusOK,"application/json", show)
}

func GetBook(ctx *gin.Context){
	bookId := ctx.Param("id")
	idData,_ := strconv.Atoi(bookId)
	state := false
	var bookData Book


	for _, book := range bookDatas{
		if idData == book.BookID{
			state = true
			bookData = book
			break
		}
	}

	if !state{
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Book with id %v not found", idData),
			})
	} else {
		ctx.JSON(http.StatusOK,bookData)
	}

	
}

func UpdateBook(ctx *gin.Context){
	bookId := ctx.Param("id")
	idData,_ := strconv.Atoi(bookId)
	state := false
	var updatedBook Book

		if err := ctx.ShouldBindJSON(&updatedBook); err != nil{
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

	
	for i, book := range bookDatas{
		if idData == book.BookID{
			state = true
			bookDatas[i] = updatedBook
			bookDatas[i].BookID = idData
			break
		}
	}

	if !state{
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Book with id %v not found", idData),
			})
	} else {
		show := []byte(`"Updated"`)
		ctx.Data(http.StatusOK,"application/json", show)
	}
	
}


func DeleteBook(ctx *gin.Context){
	bookId := ctx.Param("id")
	idData,_ := strconv.Atoi(bookId)
	state := false
	var indexBook int


	
	for i, book := range bookDatas{
		if idData == book.BookID{
			state = true
			indexBook = i
			break
		}
	}

	if !state{
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Book with id %v not found", idData),
			})
		return
	}



	if state {
		copy(bookDatas[indexBook:], bookDatas[indexBook+1:])
		bookDatas[len(bookDatas)-1] = Book{}
		bookDatas = bookDatas[:len(bookDatas)-1]

		show := []byte(`"Deleted"`)
		ctx.Data(http.StatusOK,"application/json", show)
	}

	}

	func GetAllBook(ctx *gin.Context){
		ctx.JSON(http.StatusOK,bookDatas)
	}
