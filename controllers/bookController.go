package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var m sync.Mutex
type Book	struct{
	BookID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Desc string `json:"desc"`

}

const (
	host = "localhost"
	port = 5432
	user = "postgres1"
	password = "postgres1"
	dbname = "db_go_sql"
)

var (
	db *sql.DB
	err error
)


var bookDatas = []Book{}

func RunDB(){
	// m.Lock()
	dbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host,port,user,password,dbname)
	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil{
		panic(err)
	}
	fmt.Println("Connected to Database")
	// m.Unlock()
}


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
