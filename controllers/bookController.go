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


func RunDB(){
	// m.Lock()
	dbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host,port,user,password,dbname)
	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil{
		panic(err)
	}
	fmt.Println("Connected to Database")
	// m.Unlock()
}


func CreateBook(ctx *gin.Context){
	var newBook Book
	var book = Book{}
	if err := ctx.ShouldBindJSON(&newBook); err != nil{
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `INSERT INTO books ("title", "author", "desc") VALUES ($1, $2, $3) Returning *`
	err = db.QueryRow(sqlStatement, &newBook.Title, &newBook.Author, &newBook.Desc).Scan(&book.BookID,&book.Title,&book.Author,&book.Desc)
	if err != nil{
		panic(err)
	}

	show := []byte(`"Created"`)
	ctx.Data(http.StatusCreated,"application/json", show)
}

func GetBook(ctx *gin.Context){
	bookId := ctx.Param("id")
	idData,_ := strconv.Atoi(bookId)
	var bookData Book

	sqlStatement := `SELECT * FROM books WHERE id = $1`
	err := db.QueryRow(sqlStatement, idData).Scan(&bookData.BookID,&bookData.Title,&bookData.Author,&bookData.Desc)

	if err != nil{
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
	var updatedBook Book
		if err := ctx.ShouldBindJSON(&updatedBook); err != nil{
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

	sqlStatement := `UPDATE books 
	SET "title" = $2, "author" = $3, "desc" = $4
	WHERE id = $1
	`
	res, err := db.Exec(sqlStatement, idData,&updatedBook.Title, &updatedBook.Author, &updatedBook.Desc)
	count, err2 := res.RowsAffected()
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Book with id %v not found", idData),
		})
	} else if (count == 0) {
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Book with id %v not found", idData),
		})
	
	} else if err2 != nil{
		panic(err)
	} else {
		show := []byte(`"Updated"`)
		ctx.Data(http.StatusOK,"application/json", show)
	}
}


func DeleteBook(ctx *gin.Context){
	bookId := ctx.Param("id")
	idData,_ := strconv.Atoi(bookId)

	sqlStatement := `DELETE FROM books 
	WHERE id = $1
	`
	res, err := db.Exec(sqlStatement, idData)
	count, err2 := res.RowsAffected()
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Book with id %v not found", idData),
		})
	} else if (count == 0) {
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"error_message": fmt.Sprintf("Book with id %v not found", idData),
		})
	} else if err2 != nil{
		panic(err)
	} else {
		show := []byte(`"Deleted"`)
		ctx.Data(http.StatusOK,"application/json", show)
	}

	}

	func GetAllBook(ctx *gin.Context){
		var bookDatas = []Book{}
		sqlStatement := `SELECT * FROM books`
		rows, err := db.Query(sqlStatement)
		if err != nil {
			panic(err)
		}

		for rows.Next(){
			var book = Book{}
			err = rows.Scan(&book.BookID,&book.Title,&book.Author,&book.Desc)
			if err != nil{
				panic(err)
			}
			bookDatas = append(bookDatas, book)
		}



		ctx.JSON(http.StatusOK,bookDatas)
	}
