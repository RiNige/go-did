package vc

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

type VerifiableCredential struct {
	Context           []string          `json:"@context"`
	Type              []string          `json:"type"`
	Issuer            string            `json:"issuer"`
	IssuanceDate      string            `json:"issuanceDate"`
	CredentialSubject map[string]string `json:"credentialSubject"`
	Proof             Proof             `json:"proof"`
}

type Proof struct {
	Type               string `json:"type"`
	Created            string `json:"created"`
	ProofPurpose       string `json:"proofPurpose"`
	VerificationMethod string `json:"verificationMethod"`
	JWS                string `json:"jws"`
}

// CreateVC 会用 Issuer 的私钥创建一个带签名的 VC
func CreateVC(issuerDID string, issuerKey *ecdsa.PrivateKey, subjectDID string, claims map[string]string) (*VerifiableCredential, error) {
	vc := &VerifiableCredential{
		Context:      []string{"https://www.w3.org/2018/credentials/v1"},
		Type:         []string{"VerifiableCredential"},
		Issuer:       issuerDID,
		IssuanceDate: time.Now().Format(time.RFC3339),
		CredentialSubject: map[string]string{
			"id": subjectDID,
		},
	}

	// 把 claims（如 name, degree）添加到 credentialSubject
	for k, v := range claims {
		vc.CredentialSubject[k] = v
	}

	// 将 VC 的 core 数据编码为 JSON（用于签名）
	vcJSON, err := json.Marshal(vc)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal VC for signing: %v", err)
	}

	// 对 VC 签名
	hash := crypto.Keccak256Hash(vcJSON)
	signature, err := crypto.Sign(hash.Bytes(), issuerKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign VC: %v", err)
	}

	// 添加签名
	vc.Proof = Proof{
		Type:               "EcdsaSecp256k1Signature2020",
		Created:            time.Now().Format(time.RFC3339),
		ProofPurpose:       "assertionMethod",
		VerificationMethod: issuerDID + "#controllerKey",
		JWS:                fmt.Sprintf("0x%x", signature),
	}

	return vc, nil
}
