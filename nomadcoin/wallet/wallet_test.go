package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"io/fs"
	"reflect"
	"testing"

)

const (
	testKey string = "307702010104205c00f04039402a6123017783f9b164715582b2d63fa3eb18667be24541292cf5a00a06082a8648ce3d030107a144034200042880dff27aba6d6da3a23766e422c1c29a81b7f00b6c1743b863c54d1ac4b81d3d485b1b25fe7a569c7ed64f3e41bf7a4df297f2bb18d6e074d916453389b016"
	testPayload string = "0006abbbf05890d47dca4b0a022f16bd4424cbcb9234f00b119b6891"
	testSig string = "45a226005966cd375a47a086684ffa5da479321da6255fa405f2cbe80254fe0e3629d1b221c050da7a6f767cf0e38b9638f2517ca703b6e850a0bb0a4fc5271b"
)

type fakeLayer struct {
	fakeHasWalletFile func() bool
}

func (f fakeLayer) hasWalletFile() bool {
	return f.fakeHasWalletFile()
}

func (fakeLayer) writeFile(name string, data []byte, perm fs.FileMode) error {
	return nil
}

func (fakeLayer) readFile(name string) ([]byte, error) {
	return x509.MarshalECPrivateKey(makeTestWallet().privateKey)
}

func TestWallet(t *testing.T) {
	t.Run("New Wallet is created", func(t *testing.T) {
		files = fakeLayer {
			fakeHasWalletFile: func() bool {return false},
		}
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance.")
		}
	})
	t.Run("Wallet is restored", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool {return true},
		}
		w = nil
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance.")
		}
	})
}

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = aFromK(key)
	return w 
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() should return a hex decoded string, got %s", s)
	}
}

func TestVerify(t *testing.T) {
	type test struct{
		input string
		ok bool
	}
	tests := []test{
		{testPayload, true},
		{"0406abbbf05890d47dca4b0a022f16bd4424cbcb9234f00b119b6891", false},
	}
	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSig, tc.input, w.Address)
		if ok != tc.ok {
			t.Error("Verify() could not verify testSignature and testPayload.")
		}
	}
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("xx")
	if err == nil { 
		t.Error("restoreBigInts should return error when payload is not hex.")
	}
}