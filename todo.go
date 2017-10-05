package main

import (
	"database/sql"

	"go-echo-vue/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {

	db := initDB("storage.db")
	migrate(db)

	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/dummy_db")

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL AUTO_INCREMENT,
		PRIMARY KEY (` + "`id`)," + `
		name VARCHAR(256)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
