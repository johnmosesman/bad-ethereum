package keys

import (
	"testing"
)

func TestPublicAddress(t *testing.T) {
	pk := "f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315"
	addr := PublicAddress(pk)

	if addr != "0x001d3f1ef827552ae1114027bd3ecf1f086ba0f9" {
		t.Fatal("Incorrect public key generated")
	}
}
