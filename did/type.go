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
	DID        string       `json:"did"`
	PrivateKey string       `json:"private_key,omitempty"`
	Document   *DIDDocument `json:"document"`
}
