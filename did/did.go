package did

import (
	"crypto/ecdsa"
	"fmt"
	"log"

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

func NewClient(portNumber string) *ethclient.Client {
	// Connect to the blockchain running at localhost
	client, err := ethclient.Dial("http://localhost:" + portNumber)
	if err != nil {
		panic(err)
	}

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
		Document:   doc,
	}, nil
}
