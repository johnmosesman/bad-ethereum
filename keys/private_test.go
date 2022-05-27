package keys

import "testing"

func TestGenPrivateKey(t *testing.T) {
	key := GenPrivateKey()

	if len(key) != 32 {
		t.Errorf("Key is %d bytes, not 32 bytes", len(key))
	}
}
