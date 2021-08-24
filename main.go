package main

import (
	"context"
	"fmt"
	"net/http"
	"offerapp/models"
	"offerapp/routes"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	fmt.Println("main started")

	router := gin.Default()

	store := cookie.NewStore([]byte(os.Getenv("TOKEN_SECRET")))
	router.Use(sessions.Sessions("mysession", store))

	router.Static("/css", "./templates/css")

	router.LoadHTMLGlob("templates/*.html")

	router.Static("/img", "./img")

	router.Use(dbMiddleware(*conn))

	router.GET("/", authMiddleWare(), routes.IndexPage)

	appGroup := router.Group("app")
	{
		appGroup.GET("/login", routes.AppUserLogin)
		appGroup.POST("/login", routes.AppConvertLogin)
	}

	usersGroup := router.Group("users")
	{
		usersGroup.POST("register", routes.UsersRegister)
		usersGroup.POST("login", routes.UsersLogin)
		usersGroup.GET("logout", authMiddleWare(), routes.Logout)
	}

	itemsGroup := router.Group("items")
	{
		fmt.Println("Items Group involved")
		itemsGroup.POST("create", authMiddleWare(), routes.ItemsCreate)
		itemsGroup.GET("index", routes.ItemsIndex)
		itemsGroup.GET("sold_by_user", authMiddleWare(), routes.ItemsForSaleByCurrentUser)
		itemsGroup.PUT("update", authMiddleWare(), routes.ItemsUpdate)
		itemsGroup.DELETE("delete", authMiddleWare(), routes.ItemsDelete)
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

func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token := fmt.Sprintf("%v", session.Get("Authorization"))
		//bearer := c.Request.Header.Get("Authorization")
		//fmt.Printf("Authorization: %v", bearer)
		//split := strings.Split(bearer, "Bearer ")
		if len(token) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
			return
		}
		//fmt.Printf("Split 1:%v 2:%v", split[0], split[1])
		////token := bearer
		//fmt.Printf("Bearer (%v) \n", token)
		isValid, userID := models.IsTokenValid(token)
		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
		} else {
			c.Set("user_id", userID)
			c.Next()
		}
	}
}
