// package did

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"fmt"
// 	"log"

// 	"github.com/RiNige/go-did/contracts"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func newPrivateKey() (*ecdsa.PrivateKey, string, error) {
// 	// Create privat & public key
// 	privateKey, err := crypto.GenerateKey()
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
// 	}
// 	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

// 	return privateKey, address, nil
// }

// func getAccountAuth(client *ethclient.Client, privatekey string) *bind.TransactOpts {
// 	privateKey, err := crypto.HexToECDSA(privatekey)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		panic("invalid key")
// 	}

// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

// 	//fetch the last use nonce of account
// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("nounce=", nonce)

// 	chainID, err := client.ChainID(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
// 	if err != nil {
// 		panic(err)
// 	}
// 	auth.GasLimit = uint64(2000000)

// 	return auth
// }

// func NewClient(portNumber string) *ethclient.Client {
// 	// Connect to the blockchain running at localhost
// 	client, err := ethclient.Dial("http://localhost:" + portNumber)
// 	if err != nil {
// 		panic(err)
// 	}
// 	blockNum, err := client.BlockNumber(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Connected to local Ethereum network at: %s, the current Block Number is %v\n", portNumber, blockNum)
// 	return client
// }

// func HandleCreateDID() (*DIDCreationResponse, error) {
// 	privateKey, address, err := newPrivateKey()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to generate private key: %v", err)
// 	}
// 	did := fmt.Sprintf("did:ethr:%s", address)

// 	methodSpec := VerificationMethod{
// 		ID:                  did + "#controllerKey",
// 		Type:                "EcdsaSecp256k1RecoveryMethod2020",
// 		Controller:          did,
// 		BlockchainAccountId: "eip155:1:" + address,
// 	}

// 	doc := &DIDDocument{
// 		Context: []string{
// 			"https://www.w3.org/ns/did/v1",
// 			"https://w3id.org/security/suites/secp256k1recovery-2020/v2",
// 		},
// 		ID:                 did,
// 		VerificationMethod: []VerificationMethod{methodSpec},
// 	}

// 	doc.Authentication = append(doc.Authentication, methodSpec.ID)
// 	doc.AssertionMethod = append(doc.AssertionMethod, methodSpec.ID)

// 	return &DIDCreationResponse{
// 		DID:        did,
// 		PrivateKey: fmt.Sprintf("%x", crypto.FromECDSA(privateKey)),
// 		Address:    address,
// 		Document:   doc,
// 	}, nil
// }

// func DeployContract(client *ethclient.Client, privateKey string) common.Address {
// 	// Create auth and deploy contract
// 	auth := getAccountAuth(client, privateKey)
// 	contractAddress, tx, _, err := contracts.DeployContracts(auth, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	receipt, err := bind.WaitMined(context.Background(), client, tx)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if receipt.Status == 0 {
// 		log.Fatal("Failed to deploy the smart contract")
// 	}
// 	fmt.Println("Successfully deployed smart contract to the blockchain!")
// 	return contractAddress
// }

// func StoreHashOnChain(did string, hash string, owner string, eth *ethclient.Client, contract *contracts.Contracts) (string, error) {
// 	// Create auth for this connection
// 	auth := getAccountAuth(eth, owner)
// 	auth.GasLimit = 500000
// 	auth.Context = context.Background()

// 	// Send the transaction
// 	tx, err := contract.SetHash(auth, did, hash)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to set the Hash: %v", err)
// 	}

// 	// Wait until the finish of transaction asyncronously
// 	go func(txHash string, eth *ethclient.Client) {
// 		// Monitor the transaction status
// 		receipt, _ := eth.TransactionReceipt(context.Background(), common.HexToHash(txHash))
// 		if receipt.Status == 0 {
// 			log.Printf("Transaction %s Failed", txHash)
// 		}
// 	}(tx.Hash().Hex(), eth)

// 	return tx.Hash().Hex(), nil
// }

// func QueryHashOnChain(did string, contract *contracts.Contracts, ethClient *ethclient.Client) (string, error) {
// 	callOpts := &bind.CallOpts{
// 		Pending: false,
// 		Context: context.Background(),
// 	}

// 	hash, err := contract.GetHash(callOpts, did)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to query hash on chain: %v", err)
// 	}
// 	return hash, nil
// }

// package did

// import (
// 	"context"
// 	"crypto/ecdsa"
// 	"fmt"
// 	"log"

// 	"github.com/RiNige/go-did/contracts"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// const ganachePrivateKey = "22599c307e9b1dd3357cce5cebf440b26e3ede715cfe82496f82edf72995402c" // 不带0x

// func newPrivateKey() (*ecdsa.PrivateKey, string, error) {
// 	privateKey, err := crypto.HexToECDSA(ganachePrivateKey)
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("cannot assert publicKey as ECDSA")
// 	}
// 	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

// 	fmt.Println("✅ Using account:", address)

// 	return privateKey, address, nil
// }

// func getAccountAuth(client *ethclient.Client, privateKeyHex string) *bind.TransactOpts {
// 	privateKey, err := crypto.HexToECDSA(privateKeyHex)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		panic("invalid key")
// 	}
// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
// 	fmt.Println("🔑 Using account:", fromAddress.Hex())

// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("nonce =", nonce)

// 	chainID, err := client.ChainID(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
// 	if err != nil {
// 		panic(err)
// 	}
// 	auth.GasLimit = uint64(2000000)

// 	return auth
// }

// func NewClient(portNumber string) *ethclient.Client {
// 	client, err := ethclient.Dial("http://localhost:" + portNumber)
// 	if err != nil {
// 		panic(err)
// 	}
// 	blockNum, err := client.BlockNumber(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Connected to local Ethereum network at: %s, Block Number: %v\n", portNumber, blockNum)
// 	return client
// }

// func HandleCreateDID() (*DIDCreationResponse, error) {
// 	//privateKey, address, err := newPrivateKey()
// 	_, address, err := newPrivateKey()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to load private key: %v", err)
// 	}
// 	did := fmt.Sprintf("did:ethr:%s", address)

// 	methodSpec := VerificationMethod{
// 		ID:                  did + "#controllerKey",
// 		Type:                "EcdsaSecp256k1RecoveryMethod2020",
// 		Controller:          did,
// 		BlockchainAccountId: "eip155:1:" + address,
// 	}

// 	doc := &DIDDocument{
// 		Context: []string{
// 			"https://www.w3.org/ns/did/v1",
// 			"https://w3id.org/security/suites/secp256k1recovery-2020/v2",
// 		},
// 		ID:                 did,
// 		VerificationMethod: []VerificationMethod{methodSpec},
// 	}
// 	doc.Authentication = append(doc.Authentication, methodSpec.ID)
// 	doc.AssertionMethod = append(doc.AssertionMethod, methodSpec.ID)

// 	return &DIDCreationResponse{
// 		DID:        did,
// 		PrivateKey: ganachePrivateKey,
// 		Address:    address,
// 		Document:   doc,
// 	}, nil
// }

// func DeployContract(client *ethclient.Client, privateKey string) common.Address {
// 	auth := getAccountAuth(client, privateKey)
// 	contractAddress, tx, _, err := contracts.DeployContracts(auth, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	receipt, err := bind.WaitMined(context.Background(), client, tx)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if receipt.Status == 0 {
// 		log.Fatal("Smart contract deployment failed")
// 	}
// 	fmt.Println("📦 Smart contract deployed to:", contractAddress.Hex())
// 	return contractAddress
// }

// func StoreHashOnChain(did string, hash string, privateKey string, eth *ethclient.Client, contract *contracts.Contracts) (string, error) {
// 	auth := getAccountAuth(eth, privateKey)
// 	auth.GasLimit = 500000
// 	auth.Context = context.Background()

// 	tx, err := contract.SetHash(auth, did, hash)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to set the Hash: %v", err)
// 	}

// 	go func(txHash string) {
// 		receipt, _ := eth.TransactionReceipt(context.Background(), common.HexToHash(txHash))
// 		if receipt != nil && receipt.Status == 0 {
// 			log.Printf("Transaction %s Failed", txHash)
// 		}
// 	}(tx.Hash().Hex())

// 	return tx.Hash().Hex(), nil
// }

package did

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/RiNige/go-did/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func newPrivateKey() (*ecdsa.PrivateKey, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return privateKey, address, nil
}

func getAccountAuth(client *ethclient.Client, privatekey string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(privatekey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nonce=", nonce)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(2000000)

	return auth
}

func NewClient(portNumber string) *ethclient.Client {
	client, err := ethclient.Dial("http://localhost:" + portNumber)
	if err != nil {
		panic(err)
	}
	blockNum, err := client.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to local Ethereum network at: %s, the current Block Number is %v\n", portNumber, blockNum)
	return client
}

func HandleCreateDID() (*DIDCreationResponse, error) {
	privateKey, address, err := newPrivateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}
	did := fmt.Sprintf("did:ethr:%s", address)

	methodSpec := VerificationMethod{
		ID:                  did + "#controllerKey",
		Type:                "EcdsaSecp256k1RecoveryMethod2020",
		Controller:          did,
		BlockchainAccountId: "eip155:1:" + address,
	}

	doc := &DIDDocument{
		Context: []string{
			"https://www.w3.org/ns/did/v1",
			"https://w3id.org/security/suites/secp256k1recovery-2020/v2",
		},
		ID:                 did,
		VerificationMethod: []VerificationMethod{methodSpec},
	}
	doc.Authentication = append(doc.Authentication, methodSpec.ID)
	doc.AssertionMethod = append(doc.AssertionMethod, methodSpec.ID)

	return &DIDCreationResponse{
		DID:        did,
		PrivateKey: fmt.Sprintf("%x", crypto.FromECDSA(privateKey)),
		Address:    address,
		Document:   doc,
	}, nil
}

func DeployContract(client *ethclient.Client, privateKey string) common.Address {
	auth := getAccountAuth(client, privateKey)
	contractAddress, tx, _, err := contracts.DeployContracts(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		panic(err)
	}
	if receipt.Status == 0 {
		log.Fatal("Failed to deploy the smart contract")
	}
	fmt.Println("Successfully deployed smart contract to the blockchain!")
	return contractAddress
}

func StoreHashOnChain(did string, hash string, owner string, eth *ethclient.Client, contract *contracts.Contracts) (string, error) {
	auth := getAccountAuth(eth, owner)
	auth.GasLimit = 500000
	auth.Context = context.Background()

	tx, err := contract.SetHash(auth, did, hash)
	if err != nil {
		return "", fmt.Errorf("failed to set the Hash: %v", err)
	}

	// Wait until the finish of transaction asyncronously
	go func(txHash string, eth *ethclient.Client) {
		receipt, _ := eth.TransactionReceipt(context.Background(), common.HexToHash(txHash))
		if receipt.Status == 0 {
			log.Printf("Transaction %s Failed", txHash)
		}
	}(tx.Hash().Hex(), eth)

	return tx.Hash().Hex(), nil
}

func QueryHashOnChain(did string, contract *contracts.Contracts, ethClient *ethclient.Client) (string, error) {
	callOpts := &bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	}
	hash, err := contract.GetHash(callOpts, did)
	if err != nil {
		return "", fmt.Errorf("failed to query hash on chain: %v", err)
	}
	return hash, nil
}
