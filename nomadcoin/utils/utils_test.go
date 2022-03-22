package utils

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	hash := "634e992c44a70db8d5ad52e5d595fc6607e6e6f5fb28e6d861012dfd"
	s := struct{Test string}{Test: "test"}
	t.Run("Hash is always same", func(t *testing.T) {
		x := Hash(s)
		if x != hash {
		t.Errorf("Expected %s, got %s", hash, x)
		}
	})
	t.Run("Hash is hex encoded", func(t *testing.T) {
		x := Hash(s)
		_, err := hex.DecodeString(x)
		if err != nil {
			t.Error("Hash should be hex encoded")
		}
	})
} 

func ExampleHash() {
	s := struct{Test string}{Test: "test"}
	x := Hash(s)
	fmt.Println(x)
	// Output: 634e992c44a70db8d5ad52e5d595fc6607e6e6f5fb28e6d861012dfd
}