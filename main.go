package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Users struct {
	ID        int    `query:"id" form:"id" json:"id"`
	FirstName string `query:"fname" form:"fname" json:"FirstName"`
	LastName  string `query:"lname" form:"lname" json:"LastName"`
	Email     string `query:"email" form:"email" json:"Email"`
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	//Echo instance
	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}

	e.GET("/", index)

	e.POST("/search", search)

	e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error {
	users := []Users{}
	return c.Render(http.StatusOK, "index.html", users)
}

func search(c echo.Context) error {

	db := dbConnect()
	//Get search string from form
	var searchString string = c.FormValue("fname")
	var rows *sql.Rows
	var err error

	rows, err = db.Query(fmt.Sprintf("SELECT fname, lname, email FROM users WHERE UPPER(users.fname) LIKE UPPER ('%s')", searchString))
	checkErr(err)

	user := Users{}
	users := []Users{}

	for rows.Next() {
		var fname, lname, email string
		err = rows.Scan(&fname, &lname, &email)
		checkErr(err)
		user.FirstName = fname
		user.LastName = lname
		user.Email = email
		users = append(users, user)
	}

	defer db.Close()

	return c.Render(http.StatusOK, "index.html", users)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//ensure that your local postgres db has the name, port, username, and password as stated below
//create a table called "users" inside that database.
func dbConnect() (db *sql.DB) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Could not load env variables, Err: %s", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	//connect to a db by creating a string that contains the db information
	pgsl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//to ensure the connection actually happens use sql.DB connection commands that are a part of the lib/pq package
	db, err = sql.Open("postgres", pgsl)
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	fmt.Fprintf(os.Stdout, "You have connected to the database successfully\n")
	return db
}
