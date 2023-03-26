package controllers

import (
	"net/http"

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


	ctx.JSON(http.StatusOK,gin.H{
	"message":"Created",
	})
}