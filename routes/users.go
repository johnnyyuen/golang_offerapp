package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"

	"offerapp/models"

	"github.com/gin-contrib/sessions"
)

func writeToSession(c *gin.Context, t *string) {
	session := sessions.Default(c)
	session.Set("Authorization", &t)
	session.Save()
}

func UsersLogin(c *gin.Context) {
	//fmt.Printf("hello user login here")
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		c.Redirect(http.StatusMovedPermanently, "/app/login")
		return
	}
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)
	err = user.IsAuthenticated(&conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
		c.Redirect(http.StatusMovedPermanently, "/app/login")
		return
	}
	token, err := user.GetAuthToken()
	if err == nil {
		writeToSession(c, &token)
		//c.JSON(http.StatusOK, gin.H{
		//	"token": token,
		//})
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

}

func UsersRegister(c *gin.Context) {

	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)
	err = user.Register(&conn)
	if err != nil {
		fmt.Println("Error in user.Register()")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := user.GetAuthToken()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": user.ID,
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign out successfully",
	})
}
