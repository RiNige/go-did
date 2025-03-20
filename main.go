package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
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

	// Initiate local Etherum connection
	eth := did.NewClient("7545")
	defer eth.Close()

	// Deploy smart contract to local blockchain
	contractAddress := did.DeployContract(eth, "27607949c7345cf1142c809afded87af7c63cc78c15061112373c8dc69952ce7")
	fmt.Println(contractAddress)

	// Initiate Gin Server
	r := gin.Default()
	r.POST("/dids", func(c *gin.Context) {
		// Generate DID
		resp, err := did.HandleCreateDID()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to create the DID Document",
				"details": err.Error(),
			})
			return
		}

		// Compute the Hash for DID Document
		respByte, err := json.Marshal(resp.Document)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert DID Document to byte"})
			return
		}

		hash := sha256.Sum256(respByte)
		hashHex := hex.EncodeToString(hash[:])
		fmt.Println(hashHex)

		// Save DID Document to Postgres on AWS
		err = dbclient.SaveDID(db.DIDRecord{
			DID:      resp.DID,
			Document: string(respByte),
			Hash:     hashHex,
			Owner:    resp.Address,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save DID Document to Postgres"})
			return
		}
		// Return the Result
		c.JSON(http.StatusOK, gin.H{
			"DID":           resp.DID,
			"StoreOn Chain": resp.StoreOnChain,
			"DBHash":        hashHex,
		})
	})
	r.Run(":8080")
}
