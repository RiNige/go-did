非常好！你给出的这段是项目的**真正可运行的 `main.go` 文件**，包含了完整的 HTTP 接口服务。下面我将 **分模块、逐段逐行讲解这段代码的结构、功能、执行逻辑和背后的原理**。

---

## 🧠 一、`main.go` 总体结构

```text
1. 初始化数据库和区块链连接
2. 加载已部署的智能合约
3. 启动 Gin HTTP 服务并注册 6 个接口：
   ├── /vc/create       → 创建 Verifiable Credential
   ├── /vc/:id          → 查询某个 VC
   ├── /vp/create       → 创建 Verifiable Presentation
   ├── /vp/verify       → 验证 VP 签名
   ├── /vc/onchain      → VC 上链存证（存储哈希）
   └── /vc/verify_onchain/:vcID → 验证 VC 是否真的上过链
```

---

## 📦 二、代码逐段详细讲解

---

### 🔹1. 引入包

```go
import (
	"crypto/sha256"  // 做 VC 哈希
	"encoding/hex"   // 转十六进制
	"encoding/json"  // JSON 序列化
	"net/http"
	"time"

	// 本地模块
	"github.com/RiNige/go-did/contracts"
	"github.com/RiNige/go-did/db"
	"github.com/RiNige/go-did/did"
	"github.com/RiNige/go-did/vc"
	"github.com/RiNige/go-did/vp"

	"github.com/ethereum/go-ethereum/common" // 地址转换
	"github.com/ethereum/go-ethereum/crypto" // 密钥生成和签名
	"github.com/gin-gonic/gin"               // Web框架
)
```

---

### 🔹2. 初始化数据库 + 区块链 + 合约实例

```go
dbClient := db.ConnectDB()
defer dbClient.Close()

ethClient := did.NewClient("7545")         // Ganache 默认端口
defer ethClient.Close()

contractAddressStr := "0x9d84..."          // 🔧 替换为实际部署的合约地址
contractAddress := common.HexToAddress(contractAddressStr)
contractClient, err := contracts.NewContracts(contractAddress, ethClient)
```

- **数据库连接**用于存储 VC
- **ethClient** 是本地以太坊节点（如 Ganache）
- **contractClient** 是 Go 合约绑定，方便后面调用合约方法

---

## 🔧 接口部分（Gin HTTP）

---

### 1️⃣ `/vc/create` — 签发 VC

```go
r.POST("/vc/create", func(c *gin.Context) {
	var req struct {
		SubjectDID string            `json:"subjectDID"`
		Claims     map[string]string `json:"claims"`
	}
```

- 前端提交：持有者 DID + claims（如 name, email）
- 示例 JSON 请求体：
```json
{
  "subjectDID": "did:ethr:0xabc...",
  "claims": {
    "name": "Alice",
    "role": "Student"
  }
}
```

```go
privateKey, _ := crypto.HexToECDSA("私钥")
issuerDID := "did:ethr:0x391..."

vcData, err := vc.CreateVC(issuerDID, privateKey, req.SubjectDID, req.Claims)
```

- 调用模块 `vc.CreateVC` 创建 VC
- 生成完整结构 + 签名

```go
vcJSON, _ := json.Marshal(vcData)
hash := sha256.Sum256(vcJSON)
vcID := hex.EncodeToString(hash[:])
```

- 将 VC 的 JSON 做哈希，生成唯一 ID（作为数据库主键）

```go
dbClient.SaveVC(...) // 存入数据库
```

---

### 2️⃣ `/vc/:id` — 查询某个 VC

```go
r.GET("/vc/:id", func(c *gin.Context) {
	vcID := c.Param("id")
	record, err := dbClient.GetVC(vcID)
```

- 从数据库查出 VC 并返回给用户
- 提供原始 VC JSON、签发者、持有者、签名等字段

---

### 3️⃣ `/vp/create` — 创建 VP

```go
r.POST("/vp/create", func(c *gin.Context) {
	var req struct {
		VCID      string   `json:"vcID"`      // 哪张 VC
		Fields    []string `json:"fields"`    // 选取哪些 claims
		HolderDID string   `json:"holderDID"` // 谁来持有并签名
	}
```

流程如下：
1. 根据 VC ID 查询数据库拿到 VC
2. 从 claims 中提取需要的字段 → 构建 VP 的 payload
3. 用 Holder 的私钥签名 → 得到 VP
4. 返回 VP 给前端

```go
holderKey := crypto.HexToECDSA("...")
vpData, err := vp.CreateVP(req.VCID, selected, req.HolderDID, holderKey)
```

---

### 4️⃣ `/vp/verify` — 验证 VP

```go
r.POST("/vp/verify", func(c *gin.Context) {
	var req struct {
		VP vp.VerifiablePresentation `json:"vp"`
	}
```

- 前端把 VP 发过来
- 后端使用 `vp.VerifyVP()` 方法
  - 拿到 VP 的 holder DID
  - 获取其公钥
  - 校验签名（内容 hash 是否和签名匹配）

---

### 5️⃣ `/vc/onchain` — VC 上链（存储哈希）

```go
r.POST("/vc/onchain", func(c *gin.Context) {
	var req struct {
		VCID       string `json:"vcID"`
		PrivateKey string `json:"privateKey"` // 发起交易需要私钥
	}
```

- 拿到 VCID → 查出该 VC 的哈希（vcID 本身就是 hash）
- 使用合约方法 `StoreHashOnChain(subjectDID, vcID, privateKey...)`
- 返回交易哈希 `txHash`

---

### 6️⃣ `/vc/verify_onchain/:vcID` — 链上查验 VC 是否存在

```go
r.GET("/vc/verify_onchain/:vcID", func(c *gin.Context) {
	vcID := c.Param("vcID")
```

- 查出本地 VC 的 `subject` DID
- 从链上查出该 DID 对应存过的哈希
- 对比是否等于本地 `vcID` → 匹配则说明“这张 VC 确实曾上链存证过”

---

### ✅ 启动服务

```go
r.Run(":8080")
```

服务运行在 `http://localhost:8080`

---

## 📌 总结：接口作用对照表

| 接口路径 | 功能 |
|----------|------|
| `POST /vc/create` | 签发 VC 并保存到数据库 |
| `GET /vc/:id` | 查询 VC 详情 |
| `POST /vp/create` | 选择性生成 VP |
| `POST /vp/verify` | 验证 VP 的签名有效性 |
| `POST /vc/onchain` | 将 VC 哈希存入智能合约 |
| `GET /vc/verify_onchain/:vcID` | 查询该 VC 是否上过链 |

---

## ✅ 推荐下一步：

你可以继续选一个方向我帮你深入讲解：

1. 🔍 拆解 `vc.CreateVC` 的代码内部（怎么签名、怎么构造 JSON）
2. 🔐 分析 VP 签名验证的原理和 `vp.VerifyVP()` 细节
3. 🔗 讲解合约 `StoreHashOnChain` / `QueryHashOnChain` 的源码逻辑
4. 💻 教你本地运行该项目并进行测试（Postman 示例）

你想看哪一部分？我继续详细拆给你！