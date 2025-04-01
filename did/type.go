package did

// Specify the DID verification method based on the doc:
// https://github.com/decentralized-identity/ethr-did-resolver/blob/master/doc/did-method-spec.md#method-specific-identifier
type VerificationMethod struct {
	ID                  string `json:"id"`
	Type                string `json:"type"`
	Controller          string `json:"controller"`
	PublicKeyHex        string `json:"publicKeyHex,omitempty"`
	BlockchainAccountId string `json:"blockchainAccountId,omitempty"`
}

// Specify the overall structure of DID Document, based on:
// https://github.com/decentralized-identity/ethr-did-resolver/blob/master/doc/did-method-spec.md#crud-operation-definitions
type DIDDocument struct {
	Context            []string             `json:"@context"`
	ID                 string               `json:"id"`
	VerificationMethod []VerificationMethod `json:"verificationMethod"`
	Authentication     []string             `json:"authentication"`
	AssertionMethod    []string             `json:"assertionMethod"`
}

// Proposed Response for DID Creation
type DIDCreationResponse struct {
	DID          string       `json:"did"`
	PrivateKey   string       `json:"private_key,omitempty"`
	Address      string       `json:"address"`
	StoreOnChain bool         `json:"storeonchain"`
	Document     *DIDDocument `json:"document"`
}

// Proposed Response for DID Verification
type VerificationResponse struct {
	DID          string `json:"did"`
	IsValid      bool   `json:"is_valid"`
	DBHash       string `json:"db_hash"`
	ChainHash    string `json:"chain_hash"`
	ChainNetwork string `json:"chain_network"`
	LastVerified string `json:"last_verified"`
}

// Proposed Response for DID Enquiry
type DIDResponse struct {
	DID          string       `json:"did"`
	Document     *DIDDocument `json:"document"`
	Hash         string       `json:"hash"`
	Owner        string       `json:"owner"`
	CreatedAt    string       `json:"created_at"`
	StoreOnChain bool         `json:"storeonchain"`
	ChainTXHash  string       `json:"chain_tx_hash,omitempty"`
}
