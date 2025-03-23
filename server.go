package main

import (
	"encoding/json"
	"net/http"

	"crypto/sha256"
	"encoding/hex"

	"github.com/RiNige/go-did/db"
	"github.com/RiNige/go-did/vc"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

var dbClient *db.PostgresDB

func handleCreateVC(c *gin.Context) {
	var req struct {
		IssuerPrivKey string            `json:"issuerPrivateKey"`
		IssuerDID     string            `json:"issuerDID"`
		SubjectDID    string            `json:"subjectDID"`
		Claims        map[string]string `json:"claims"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	privKey, err := crypto.HexToECDSA(req.IssuerPrivKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid private key"})
		return
	}

	vcData, err := vc.CreateVC(req.IssuerDID, privKey, req.SubjectDID, req.Claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create VC", "details": err.Error()})
		return
	}

	// 生成唯一 VC ID
	vcJSON, _ := json.Marshal(vcData)
	hash := sha256.Sum256(vcJSON)
	vcID := hex.EncodeToString(hash[:])

	err = dbClient.SaveVC(db.VCRecord{
		ID:           vcID,
		Issuer:       vcData.Issuer,
		Subject:      vcData.CredentialSubject["id"],
		Claims:       vcData.CredentialSubject,
		IssuanceDate: vcData.IssuanceDate,
		Signature:    vcData.Proof.JWS,
		Raw:          vcJSON,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save VC to DB", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "✅ VC created successfully",
		"id":      vcID,
		"vc":      vcData,
	})
}
