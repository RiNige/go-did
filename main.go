package main

import (
	"net/http"

	"github.com/RiNige/go-did/db"
	"github.com/RiNige/go-did/did"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initiate and test out PostgresSQL connection
	// NOTE: HERE WE ASSUMES THE LOCAL HOST ALREADY SETUP WITH AWS CREDENTIALS
	//ctx := context.Background()
	dbclient := db.ConnectDB()
	defer dbclient.Close()
	err := dbclient.Ping()
	if err != nil {
		panic(err)
	}

	// Initiate local Etherum connection
	eth := did.NewClient("7545")
	defer eth.Close()

	// Initiate Gin Server
	r := gin.Default()
	r.POST("/dids", func(c *gin.Context) {
		// Generate DID
		resp, err := did.HandleCreateDID()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "DID生成失败",
				"details": err.Error(),
			})
			return
		}

		// Return the Result
		c.JSON(http.StatusOK, resp)
	})
	r.Run(":8080")
}
