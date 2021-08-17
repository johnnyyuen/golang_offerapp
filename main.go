package main

import (
	"context"
	"fmt"
	"offerapp/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	//"golang.org/x/net/context"
)

func getConnectString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), "localhost", "5432", os.Getenv("DBNAME"))
}

func main() {

	conn, err := connectDB()
	if err != nil {
		return
	}

	router := gin.Default()

	router.Use(dbMiddleware(*conn))

	usersGroup := router.Group("users")
	{
		usersGroup.POST("register", routes.UsersRegister)
		usersGroup.POST("login", routes.UsersLogin)
	}

	router.Run(":3000")

}

func connectDB() (c *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(), getConnectString())
	if err != nil {
		fmt.Println("Error connecting to DB")
		fmt.Println(err.Error())
	}
	_ = conn.Ping(context.Background())
	return conn, err

}

func dbMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		c.Set("db", conn)
		c.Next()
	}
}
