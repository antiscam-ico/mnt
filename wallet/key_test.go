package wallet

import (
	"encoding/hex"
	"testing"
)

func TestSeed(t *testing.T) {
	mnemonic := "suffer draft bacon typical start retire air sniff large biology mail diagram"
	seed, err := Seed(mnemonic)
	if err != nil {
		t.Fatal(err)
	}

	if hex.EncodeToString(seed) != validSeed {
		t.Fatalf("seed got %s, want %s", hex.EncodeToString(seed), validSeed)
	}
}

func TestPublicKeyByPrivateKey(t *testing.T) {
	pubKey, err := PublicKeyByPrivateKey(validPrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	if pubKey != validPublicKey {
		t.Fatalf("PublicKey got %s, want %s", pubKey, validPublicKey)
	}
}

func TestAddressByPublicKey(t *testing.T) {
	address, err := AddressByPublicKey(validPublicKey)
	if err != nil {
		t.Fatal(err)
	}

	if address != validAddress {
		t.Fatalf("Address got %s, want %s", address, validAddress)
	}
}
