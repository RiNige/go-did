// package main

// import (
// 	"crypto/sha256"
// 	"database/sql"
// 	"encoding/hex"
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/RiNige/go-did/contracts"
// 	"github.com/RiNige/go-did/db"
// 	"github.com/RiNige/go-did/did"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/gin-gonic/gin"
// )

// // Some global parameters
// var (
// 	dbClient       *sql.DB
// 	ethClient      *ethclient.Client
// 	contractClient *contracts.Contracts
// )

// func main() {
// 	// Initiate and test out PostgresSQL connection
// 	// NOTE: HERE WE ASSUMES THE LOCAL HOST ALREADY SETUP WITH AWS CREDENTIALS
// 	//ctx := context.Background()
// 	dbClient := db.ConnectDB()
// 	defer dbClient.Close()

// 	// Initiate local Etherum connection
// 	ethClient := did.NewClient("7545")
// 	defer ethClient.Close()

// 	// Deploy smart contract to local blockchain
// 	contractAddress := did.DeployContract(ethClient, "27607949c7345cf1142c809afded87af7c63cc78c15061112373c8dc69952ce7")
// 	contractClient, err := contracts.NewContracts(contractAddress, ethClient)
// 	if err != nil {
// 		log.Fatalf("Failed to create the smart contract instance:%v", err)
// 	}

// 	// Initiate Gin Server
// 	r := gin.Default()
// 	r.POST("/dids", func(c *gin.Context) {

// 		// Generate DID
// 		resp, err := did.HandleCreateDID()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error":   "Failed to create the DID Document",
// 				"details": err.Error(),
// 			})
// 			return
// 		}

// 		// Compute the Hash for DID Document
// 		respByte, err := json.Marshal(resp.Document)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert DID Document to byte"})
// 			return
// 		}

// 		hash := sha256.Sum256(respByte)
// 		hashHex := hex.EncodeToString(hash[:])

// 		// Save DID Document to Postgres on AWS
// 		err = dbClient.SaveDID(db.DIDRecord{
// 			DID:      resp.DID,
// 			Document: string(respByte),
// 			Hash:     hashHex,
// 			Owner:    resp.Address,
// 		})
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error":       "Failed to save DID Document to Postgres",
// 				"Description": err.Error()})
// 			return
// 		}

// 		// Store the DID Document Hash onto Blockchain
// 		tx, err := did.StoreHashOnChain(resp.DID, hashHex, resp.PrivateKey, ethClient, contractClient)

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error":       "Failed to save DID Document to Blockchain",
// 				"Description": err.Error()})
// 			return
// 		}

// 		// Return the Result
// 		c.JSON(http.StatusOK, gin.H{
// 			"DID":       resp.DID,
// 			"ChainHash": tx,
// 			"DBHash":    hashHex,
// 		})
// 	})
// 	r.Run(":8080")
// }

// 执行一整套 DID ➝ VC ➝ VP ➝ Verify 的逻辑，输出到终端。从101行到192行是没写接口前的测试
// package main

// import (
// 	"crypto/sha256"
// 	"encoding/hex"
// 	"encoding/json"
// 	"fmt"
// 	"log"

// 	"github.com/RiNige/go-did/db"
// 	"github.com/RiNige/go-did/vc"
// 	"github.com/RiNige/go-did/vp" // ✅ 新增：引入 VP 包
// 	"github.com/ethereum/go-ethereum/crypto"
// )

// func main() {
// 	// ✅ 连接数据库
// 	dbClient := db.ConnectDB()
// 	defer dbClient.Close()

// 	// 👤 使用 Ganache 的第一个账户作为 Issuer 和 VP 持有者
// 	privateKey, err := crypto.HexToECDSA("22599c307e9b1dd3357cce5cebf440b26e3ede715cfe82496f82edf72995402c")
// 	if err != nil {
// 		panic(err)
// 	}

// 	issuerDID := "did:ethr:0x391358442FcEd907789Ab02899846d1Fd65BCb1E"
// 	subjectDID := issuerDID // ✅ 保证 VP 持有者的私钥和 DID 匹配
// 	//subjectDID := "did:ethr:0xabcabcabcabcabcabcabcabcabcabcabcabcabca" // 测试用 Subject DID

// 	claims := map[string]string{
// 		"name":   "Alice",
// 		"degree": "Master of Blockchain",
// 	}

// 	// 🧠 创建 VC
// 	vcData, err := vc.CreateVC(issuerDID, privateKey, subjectDID, claims)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 📄 打印 VC JSON
// 	vcJSON, _ := json.MarshalIndent(vcData, "", "  ")
// 	fmt.Println(string(vcJSON))

// 	// 🧩 生成唯一 ID（用 VC JSON 的哈希）
// 	hash := sha256.Sum256(vcJSON)
// 	vcID := hex.EncodeToString(hash[:])

// 	// 💾 保存 VC 到数据库
// 	err = dbClient.SaveVC(db.VCRecord{
// 		ID:           vcID,
// 		Issuer:       vcData.Issuer,
// 		Subject:      vcData.CredentialSubject["id"],
// 		Claims:       vcData.CredentialSubject,
// 		IssuanceDate: vcData.IssuanceDate,
// 		Signature:    vcData.Proof.JWS,
// 		Raw:          vcJSON,
// 	})

// 	if err != nil {
// 		log.Fatal("❌ Failed to save VC to DB:", err)
// 	}

// 	fmt.Println("✅ VC 已成功保存到数据库 ✅")

// 	// ✅✅✅ 新增部分：基于 VC 创建 VP（选择部分字段）
// 	selectedClaims := map[string]string{
// 		"name": vcData.CredentialSubject["name"],
// 	}

// 	vpData, err := vp.CreateVP(vcID, selectedClaims, subjectDID, privateKey)
// 	if err != nil {
// 		log.Fatal("❌ Failed to create VP:", err)
// 	}

// 	vpJSON, _ := json.MarshalIndent(vpData, "", "  ")
// 	fmt.Println("🧾 Verifiable Presentation:")
// 	fmt.Println(string(vpJSON))

// 	// ✅✅✅✅ 新增部分：验证 VP 签名是否有效
// 	ok, err := vp.VerifyVP(*vpData)
// 	if err != nil {
// 		log.Fatal("❌ VP 签名验证失败:", err)
// 	}
// 	if ok {
// 		fmt.Println("✅ VP 签名验证通过 ✅")
// 	} else {
// 		fmt.Println("❌ VP 签名验证失败 ❌")
// 	}
// }

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/RiNige/go-did/contracts"
	"github.com/RiNige/go-did/db"
	"github.com/RiNige/go-did/did"
	"github.com/RiNige/go-did/vc"
	"github.com/RiNige/go-did/vp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

func main() {
	dbClient := db.ConnectDB()
	defer dbClient.Close()

	// 🔧 MODIFY HERE: Ganache 端口（默认是7545，如果别人改了端口，要改这里）
	ethClient := did.NewClient("7545")
	defer ethClient.Close()

	// 🔧 MODIFY HERE: 智能合约部署后的地址（每台电脑部署后地址都不一样）
	contractAddressStr := "0x9d84964766677c5c62ab65e4c1c862cd1c5efe15"
	contractAddress := common.HexToAddress(contractAddressStr)

	contractClient, err := contracts.NewContracts(contractAddress, ethClient)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// ─── VC 签发 ───
	r.POST("/vc/create", func(c *gin.Context) {
		var req struct {
			SubjectDID string            `json:"subjectDID"`
			Claims     map[string]string `json:"claims"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		// 🔧 MODIFY HERE: issuer 私钥（必须是当前 Ganache 中的某个账户的私钥）
		privateKey, _ := crypto.HexToECDSA("22599c307e9b1dd3357cce5cebf440b26e3ede715cfe82496f82edf72995402c")

		// 🔧 MODIFY HERE: issuer DID（必须与上面私钥对应的账户一致）
		issuerDID := "did:ethr:0x391358442FcEd907789Ab02899846d1Fd65BCb1E"

		vcData, err := vc.CreateVC(issuerDID, privateKey, req.SubjectDID, req.Claims)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create VC"})
			return
		}

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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store VC to DB"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"vcID":      vcID,
			"vc":        vcData,
			"message":   "VC created and stored successfully",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// ─── VC 查询 ───
	r.GET("/vc/:id", func(c *gin.Context) {
		vcID := c.Param("id")
		record, err := dbClient.GetVC(vcID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "VC not found", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"vcID":      record.ID,
			"issuer":    record.Issuer,
			"subject":   record.Subject,
			"claims":    record.Claims,
			"signature": record.Signature,
			"issuedAt":  record.IssuanceDate,
			"raw":       json.RawMessage(record.Raw),
		})
	})

	// ─── VP 创建 ───
	r.POST("/vp/create", func(c *gin.Context) {
		var req struct {
			VCID      string   `json:"vcID"`
			Fields    []string `json:"fields"`
			HolderDID string   `json:"holderDID"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format", "details": err.Error()})
			return
		}

		vcRecord, err := dbClient.GetVC(req.VCID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "VC not found", "details": err.Error()})
			return
		}

		selected := map[string]string{}
		for _, key := range req.Fields {
			if val, ok := vcRecord.Claims[key]; ok {
				selected[key] = val
			}
		}

		// 🔧 MODIFY HERE: Holder 的私钥（建议和 issuer 用同一个账户测试）
		holderKey, _ := crypto.HexToECDSA("22599c307e9b1dd3357cce5cebf440b26e3ede715cfe82496f82edf72995402c")

		vpData, err := vp.CreateVP(req.VCID, selected, req.HolderDID, holderKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create VP", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"vp":        vpData,
			"message":   "VP created successfully",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// ─── VP 验证 ───
	r.POST("/vp/verify", func(c *gin.Context) {
		var req struct {
			VP vp.VerifiablePresentation `json:"vp"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid VP format", "details": err.Error()})
			return
		}

		ok, err := vp.VerifyVP(req.VP)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Verification failed", "details": err.Error()})
			return
		}

		if ok {
			c.JSON(http.StatusOK, gin.H{
				"message":   "✅ VP signature is valid",
				"timestamp": time.Now().Format(time.RFC3339),
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "❌ VP signature is invalid"})
		}
	})

	// ─── VC 上链 ───
	r.POST("/vc/onchain", func(c *gin.Context) {
		var req struct {
			VCID       string `json:"vcID"`
			PrivateKey string `json:"privateKey"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		record, err := dbClient.GetVC(req.VCID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "VC not found", "details": err.Error()})
			return
		}

		txHash, err := did.StoreHashOnChain(record.Subject, record.ID, req.PrivateKey, ethClient, contractClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store hash on chain", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"txHash":    txHash,
			"message":   "VC hash stored on chain successfully",
			"vcID":      record.ID,
			"subject":   record.Subject,
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// ─── 链上验证 VC ───
	r.GET("/vc/verify_onchain/:vcID", func(c *gin.Context) {
		vcID := c.Param("vcID")
		record, err := dbClient.GetVC(vcID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "VC not found", "details": err.Error()})
			return
		}

		onchainHash, err := did.QueryHashOnChain(record.Subject, contractClient, ethClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query on-chain hash", "details": err.Error()})
			return
		}

		if onchainHash == vcID {
			c.JSON(http.StatusOK, gin.H{
				"message":   "✅ VC hash verified on chain",
				"vcID":      vcID,
				"onchain":   true,
				"timestamp": time.Now().Format(time.RFC3339),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message":   "❌ VC hash not found or mismatch on chain",
				"vcID":      vcID,
				"onchain":   false,
				"timestamp": time.Now().Format(time.RFC3339),
			})
		}
	})

	// ✅ 启动服务
	r.Run(":8080")
}
