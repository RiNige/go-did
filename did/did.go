package did

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/RiNige/go-did/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func newPrivateKey() (*ecdsa.PrivateKey, string, error) {
	// Create privat & public key
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

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nounce=", nonce)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.GasLimit = uint64(2000000)

	return auth
}

func NewClient(portNumber string) *ethclient.Client {
	// Connect to the blockchain running at localhost
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
	// Create auth and deploy contract
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
	// Create auth for this connection
	auth := getAccountAuth(eth, owner)
	auth.GasLimit = 500000
	auth.Context = context.Background()

	// Send the transaction
	tx, err := contract.SetHash(auth, did, hash)
	if err != nil {
		return "", fmt.Errorf("failed to set the Hash: %v", err)
	}

	// Wait until the finish of transaction asyncronously
	go func(txHash string, eth *ethclient.Client) {
		// Monitor the transaction status
		receipt, _ := eth.TransactionReceipt(context.Background(), common.HexToHash(txHash))
		if receipt.Status == 0 {
			log.Printf("Transaction %s Failed", txHash)
		}
	}(tx.Hash().Hex(), eth)

	return tx.Hash().Hex(), nil
}

func GetHashFromChain(did string, eth *ethclient.Client, contract *contracts.Contracts) (string, error) {
	hash, err := contract.GetHash(nil, did)
	if err != nil {
		return "", fmt.Errorf("failed to enquiry DID on Blockchain: %v", err)
	}

	return hash, nil
}
