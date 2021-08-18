package routes

import (
	"fmt"
	"net/http"
	"offerapp/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func ItemsIndex(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	items, err := models.GetAllItems(&conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//c.JSON(http.StatusOK, gin.H{"items": items})

	data := gin.H{
		"title": "Item List",
		"items": items,
	}

	c.HTML(http.StatusOK, "index.html", data)
}

func ItemsForSaleByCurrentUser(c *gin.Context) {
	userID := c.GetString("user_id")
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	items, err := models.GetItemsBeingSoldByUser(userID, &conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func ItemsCreate(c *gin.Context) {
	userID := c.GetString("user_id")
	fmt.Printf("userID::: %v", userID)
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	item := models.Item{}
	c.ShouldBindJSON(&item)
	err := item.Create(&conn, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func ItemsDelete(c *gin.Context) {
	userID := c.GetString("user_id")
	fmt.Printf("userID::: %v", userID)
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	itemSent := models.Item{}
	err := c.ShouldBindJSON(&itemSent)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format sent"})
		return
	}

	itemTobeDeleted, err := models.FindItemById(itemSent.ID, &conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = itemTobeDeleted.DeleteItem(&conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Item Deleted")
}

func ItemsUpdate(c *gin.Context) {
	userID := c.GetString("user_id")
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	itemSent := models.Item{}
	err := c.ShouldBindJSON(&itemSent)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format sent"})
		return
	}

	itemBeingUpdate, err := models.FindItemById(itemSent.ID, &conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if itemBeingUpdate.SellerID.String() != userID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are not authorized to update this item"})
		return
	}

	err = itemSent.Update(&conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"item": itemSent})
}
