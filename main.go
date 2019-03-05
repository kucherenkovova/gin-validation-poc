package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Listing struct {
	Name        string `json:"name" binding:"required,min=5,max=30,alpha"`
	Description string `json:"description" binding:"nefield=Name"`
	Email       string `json:"email" binding:"required,email"`
	URL         string `json:"url" binding:"omitempty,uri"`
}

func HandleCreateListing(c *gin.Context) {
	listing := &Listing{}
	if err := c.ShouldBindJSON(listing); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": fmt.Sprintf("created %s", listing.Name)})
}

func main() {
	route := gin.Default()
	route.POST("/listing", HandleCreateListing)
	if err := route.Run(":7777"); err != nil {
		log.Fatalf("you suck: %+v", err)
	}
}
