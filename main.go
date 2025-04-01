package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"

	"github.com/RiNige/go-did/contracts"
	"github.com/RiNige/go-did/db"
	"github.com/RiNige/go-did/did"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// Some global parameters
const (
	systemAccount = "27607949c7345cf1142c809afded87af7c63cc78c15061112373c8dc69952ce7"
)

var (
	dbClient       *sql.DB
	ethClient      *ethclient.Client
	contractClient *contracts.Contracts
)

func main() {
	// Initiate and test out PostgresSQL connection
	// NOTE: HERE WE ASSUMES THE LOCAL HOST ALREADY SETUP WITH AWS CREDENTIALS
	//ctx := context.Background()
	dbClient := db.ConnectDB()
	defer dbClient.Close()

	// Initiate local Etherum connection
	ethClient := did.NewClient("8545")
	defer ethClient.Close()

	// Deploy smart contract to local blockchain
	contractAddress := did.DeployContract(ethClient, systemAccount)
	contractClient, err := contracts.NewContracts(contractAddress, ethClient)
	if err != nil {
		log.Fatalf("Failed to create the smart contract instance:%v", err)
	}

	// Initiate Gin Server
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

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

		// Save DID Document to Postgres on AWS
		err = dbClient.SaveDID(db.DIDRecord{
			DID:      resp.DID,
			Document: string(respByte),
			Hash:     hashHex,
			Owner:    resp.Address,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":       "Failed to save DID Document to Postgres",
				"Description": err.Error()})
			return
		}

		// Store the DID Document Hash onto Blockchain
		tx, err := did.StoreHashOnChain(resp.DID, hashHex, systemAccount, ethClient, contractClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":       "Failed to save DID Document to Blockchain",
				"Description": err.Error()})
			return
		}

		// Return the Result
		c.JSON(http.StatusOK, gin.H{
			"DID":       resp.DID,
			"ChainHash": tx,
			"DBHash":    hashHex,
		})
	})

	// Single query of DID
	r.GET("/dids/:did", func(c *gin.Context) {
		did_target := c.Param("did")

		// Find DID Document within PG
		record, err := dbClient.GetDID(did_target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":       "Cannot find the corresponding DID on DB",
				"Description": err.Error()})
			return
		}

		// Find DID Document within Blockchain
		resp_did, err := did.GetHashFromChain(did_target, ethClient, contractClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":       "Cannot find the corresponding DID on Blockchain",
				"Description": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"DID":                record.DID,
			"Owner":              record.Owner,
			"Hash on DB":         record.Hash,
			"Hash on Blockchain": resp_did,
			"Created At":         record.CreatedAt,
		})
	})

	// Batch query of DIDs
	r.GET("/dids", func(c *gin.Context) {
		did_targets := c.QueryArray("did")
		if len(did_targets) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":       "Missing DID parameter in input",
				"description": "Nil target",
			})
			return
		}

		records, err := dbClient.GetBatchDID(did_targets)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":       "Cannot find the corresponding DID on DB",
				"Description": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"count": len(records),
			"data":  records,
		})
	})

	r.Run(":8080")
}
