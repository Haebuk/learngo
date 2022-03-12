package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/learngo/nomadcoin/utils"
)

const (
	signature string = "176692e4212123345493a148db1c04bb7447b448198fbdc8f59c0c8151567c50d3b0f8723b2fef88c144657293bde538f48d17ecf7a62b15be715ba28b1e0dea%"
	hashedMessage string = "78109dc311bd35b7c9c0bae3938bf6779d3398ddc8968b7dc38a6687"
	privateKey string = "30770201010420d7388bd93ebc0fb8cadef6166ce4aa31d9f55e664a6449b3c5004ec26e7a9abba00a06082a8648ce3d030107a14403420004d31f878f5a05b7dadf61e6bb003849fdeba44bd58d17890a0e2a479505d83a23b77a97ce2640efe32d2408602cb58cc1df139ff6076315e9cd6c574cc7aa8390"
)

func Start() {

}