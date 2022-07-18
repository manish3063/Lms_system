package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	// TODO fill this in directly or through environment variable
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	DB_DSN = "postgres://postgres:manish@localhost:5432/lms?sslmode=disable"
)

var (
	DB *sql.DB
)

type User struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email" binding:"email"`
	Password   string `json:"password" binding:"required,min=8,alphanum"`
	Type       int    `json:"type"`
}

//order structs
type Order struct {
	ID                 int    `json:"id"`
	Request_id         int    `json:"request_id"`
	Book_id            string `json:"book_id"`
	Student_id         int    `json:"student_id"`
	Issue_date         string `json:"issue_date"`
	Return_date        string `json:"return_date"`
	Actual_return_date string `json:"actual_return_date"`
	Fine               string `json:"fine"`
	Approved           string `json:"approved"`
}

type Book struct {
	Book_id          string `json:"book_id"`
	Current_stock    int    `json:"current_stock"`
	Book_name        string `json:"book_name"`
	Book_author      string `json:"book_author"`
	Book_cover_image string `json:"book_cover_image"`
}

//main function

func setupRoutes(r *gin.Engine) {

	r.POST("signup/admin", Signup)
	r.POST("/signup/user", Signup)
	r.POST("login", login)
	r.POST("logout", logout)
	r.POST("/order", isLogin(), orderRequest)

	//admin
	r.POST("/order/approve", isAdmin(), Order_approve)
	r.POST("/order/return", isLogin(), Order_return)
	r.POST("/return/approve", isAdmin(), ReturnAccept)
}

func createDBConnection() {
	var err error
	DB, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	} else {
		fmt.Println("Connected to database")

	}
	// defer DB.Close()
}

//
