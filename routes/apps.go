package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
	session := sessions.Default(c)
	userEmail := fmt.Sprintf("%v", session.Get("UserEmail"))
	data := gin.H{
		"title":     "Hello!",
		"customCSS": "index.css",
		"useremail": userEmail,
	}

	c.HTML(http.StatusOK, "index.html", data)
}
func AppUserLogin(c *gin.Context) {
	/*
		check if user is login
		if not login yet, show the login page
		if already login, redirect to all item list
	*/
	data := gin.H{
		"title":     "Please Login",
		"customCSS": "signin.css",
	}

	c.HTML(http.StatusOK, "login.html", data)

}

func AppConvertLogin(c *gin.Context) {

	login, _ := json.Marshal(map[string]string{
		"email":    c.PostForm("Email"),
		"password": c.PostForm("Password"),
	})

	responseBody := bytes.NewBuffer(login)
	fmt.Printf("hello %v", responseBody)
	request, err := http.NewRequest("POST", "/users/login", responseBody)
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := http.Client{}
	client.Do(request)

	/*defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	*/
	/*
		resp, err := http.Post("/users/login", "application/json", responseBody)
		if err != nil || resp == nil {
			fmt.Printf("%v", err.Error())
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%v", err.Error())
		}
		sb := string(body)
		fmt.Println(sb)
	*/
}
