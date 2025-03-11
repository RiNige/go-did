package main

import "github.com/nuts-foundation/go-did/did"

func main() {
	didID, err := did.ParseDID("did:example:123")
	// Empty did document:
	doc := &did.Document{
		Context: []did.URI{did.DIDContextV1URI()},
		ID:      *didID,
	}

}
