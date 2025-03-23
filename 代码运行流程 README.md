好的！以下是你基于 PostgreSQL 的完整运行流程文档，适用于你将项目分享给同学后，他们可以独立运行、测试并验证 DID + VC + VP 项目的功能👇

---

# 🧪 DID + VC + VP 项目完整运行流程指南（含 PostgreSQL 配置）

本项目基于以太坊本地私链和 PostgreSQL，实现了：

✅ DID 生成  
✅ VC（Verifiable Credential）签发和存储  
✅ VP（Verifiable Presentation）创建和验证  
✅ VC 上链存证与链上验证  

---

## ✅ Step 1️⃣ 准备环境

确保同学已安装以下依赖：

- ✅ Go（建议 Go 1.19+）
- ✅ Node.js（建议 v16+）
- ✅ Ganache CLI（用于本地以太坊网络）
- ✅ solc（Solidity 编译器）：  
  ```bash
  npm install -g solc
  ```
- ✅ abigen（Go 合约绑定工具）：  
  安装 [Go-Ethereum](https://geth.ethereum.org/downloads) 后，abigen 会包含在内。
- ✅ PostgreSQL（建议通过 pgAdmin 或命令行创建表）

---

## 🛠️ Step 2️⃣ 启动 Ganache

```bash
ganache -p 7545
```

记录第一个账户的：

- 地址，如 `0x391358442FcEd...`
- 私钥，如 `0x22599c307e9b1dd3...`

---

## 🧱 Step 3️⃣ 编译 & 绑定智能合约

进入合约目录：

```bash
cd smartContract
```

编译合约并生成 ABI + BIN 文件：

```bash
solc --abi --bin DIDHashRegistry.sol -o ../build --overwrite
```

使用 abigen 生成 Go 合约绑定代码：

```bash
abigen --abi ../build/DIDHashRegistry.abi --bin ../build/DIDHashRegistry.bin --pkg contracts --out ../contracts/DIDHashRegistry.go
```

---

## 🔗 Step 4️⃣ 修改合约地址和端口

打开 `main.go`，注意有 🔧 MODIFY HERE 的地方，根据自己 ganache 的信息修改：：

```go
// ✅ 修改为你的 Ganache 端口（通常是 7545）
ethClient := did.NewClient("7545")

// ✅ 修改为你 Ganache 部署合约后生成的地址（Contract created: ...）
contractAddressStr := "0x你的合约地址"
```

---

## 🗃️ Step 5️⃣ 初始化数据库（PostgreSQL）

打开 pgAdmin 或 PostgreSQL 命令行工具，运行以下 SQL（可一行运行）：

```sql
CREATE TABLE did_documents (did TEXT PRIMARY KEY, document TEXT, hash TEXT, owner TEXT, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
CREATE TABLE verifiable_credentials (id TEXT PRIMARY KEY, issuer TEXT, subject TEXT, claims JSONB, issuance_date TIMESTAMP, signature TEXT, raw JSONB);
```

---

## 🛠️ Step 6️⃣ 配置数据库连接

打开 `db/db.go`，找到连接字符串部分：

```go
const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "你的PostgreSQL密码"
  dbname   = "postgres"
)
```

根据你本地的 PostgreSQL 修改用户、密码、端口和数据库名。

---

## 📦 Step 7️⃣ 安装 Go 依赖

在项目根目录运行：

```bash
go mod tidy
```

---

## 🚀 Step 8️⃣ 启动项目

```bash
go run main.go
```

此时后端服务将在 `http://localhost:8080` 启动。

---

## ✅ Step 9️⃣ 测试接口，建议问ChatGpt生成命令行

你可以使用 Postman 或 curl 测试：

### 📄 创建 VC：

```bash
curl -X POST http://localhost:8080/vc/create -H "Content-Type: application/json" -d "{\"subjectDID\": \"did:ethr:0xabcabcabcabcabcabcabcabcabcabcabcabcabca\", \"claims\": {\"name\": \"Alice\", \"degree\": \"Master of Blockchain\"}}"
```

### 🪪 创建 VP：

```bash
curl -X POST http://localhost:8080/vp/create -H "Content-Type: application/json" -d "{\"vcID\": \"<替换为上一步返回的vcID>\", \"fields\": [\"name\"], \"holderDID\": \"did:ethr:<Ganache账户地址>\"}"
```

### ✅ 验证 VP：

```bash
curl -X POST http://localhost:8080/vp/verify -H "Content-Type: application/json" -d "{\"vp\": {<替换为上一步返回的完整VP对象>}}"
```

### 🔗 VC 上链：

```bash
curl -X POST http://localhost:8080/vc/onchain -H "Content-Type: application/json" -d "{\"vcID\": \"<vcID>\", \"privateKey\": \"<Ganache私钥>\"}"
```

### 🔍 验证链上 VC：

```bash
curl http://localhost:8080/vc/verify_onchain/<vcID>
```

---

## 🧠 常见问题

| 问题 | 原因 | 解决方法 |
|------|------|----------|
| `invalid opcode` | 合约地址错误或合约未部署 | 请重新部署合约并更新地址 |
| `VC not found` | 数据库没保存成功 | 检查是否正确创建表 & 保存 |
| `signature invalid` | VP 签名不匹配 | 检查私钥、holderDID 与 VC issuer 是否一致 |

---
