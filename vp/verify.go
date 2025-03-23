package vp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// ✅ 验证 VP 的签名
func VerifyVP(vpData VerifiablePresentation) (bool, error) {
	// Step 1: 还原出签名内容
	claimsJSON, err := json.Marshal(vpData.VerifiableCredential)
	if err != nil {
		return false, fmt.Errorf("failed to marshal claims for verification: %v", err)
	}
	hash := crypto.Keccak256Hash(claimsJSON)

	// Step 2: 提取签名值
	sig := common.FromHex(vpData.Proof.JWS)
	if len(sig) != 65 {
		return false, errors.New("invalid signature length")
	}

	// ❌ 不再需要修复 V 值
	// sig[64] -= 27  <-- 删除这一行
	// sig[64] -= 27 // 修复以太坊签名格式

	// Step 3: 恢复公钥
	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return false, fmt.Errorf("failed to recover pubkey from signature: %v", err)
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey).Hex()

	// Step 4: 对比 DID 地址是否匹配
	didAddr := extractDIDAddress(vpData.Proof.VerificationMethod)
	if didAddr != recoveredAddr {
		return false, fmt.Errorf("signature not match with DID: got %s, expect %s", recoveredAddr, didAddr)
	}

	return true, nil
}

func extractDIDAddress(verificationMethod string) string {
	start := len("did:ethr:")
	end := len(verificationMethod)
	if hashIdx := len(verificationMethod) - len("#controllerKey"); hashIdx > start {
		end = hashIdx
	}
	return verificationMethod[start:end]
}
