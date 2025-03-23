package vp

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

type VerifiablePresentation struct {
	Type                 string              `json:"type"`
	VerifiableCredential PresentedCredential `json:"verifiableCredential"`
	Proof                Proof               `json:"proof"`
}

type PresentedCredential struct {
	ID     string            `json:"id"`
	Claims map[string]string `json:"claims"` // 选择要暴露的字段，如只展示 name
}

type Proof struct {
	Type               string `json:"type"`
	Created            string `json:"created"`
	ProofPurpose       string `json:"proofPurpose"`
	VerificationMethod string `json:"verificationMethod"`
	JWS                string `json:"jws"`
}

// ✅ 创建 VP
func CreateVP(vcID string, claims map[string]string, holderDID string, holderKey *ecdsa.PrivateKey) (*VerifiablePresentation, error) {
	vp := &VerifiablePresentation{
		Type: "VerifiablePresentation",
		VerifiableCredential: PresentedCredential{
			ID:     vcID,
			Claims: claims,
		},
	}

	vpJSON, err := json.Marshal(vp.VerifiableCredential)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal VC for VP: %v", err)
	}

	hash := crypto.Keccak256Hash(vpJSON)
	signature, err := crypto.Sign(hash.Bytes(), holderKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign VP: %v", err)
	}

	vp.Proof = Proof{
		Type:               "EcdsaSecp256k1Signature2020",
		Created:            time.Now().Format(time.RFC3339),
		ProofPurpose:       "authentication",
		VerificationMethod: holderDID + "#controllerKey",
		JWS:                fmt.Sprintf("0x%x", signature),
	}

	return vp, nil
}
