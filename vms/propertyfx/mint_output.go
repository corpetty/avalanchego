package propertyfx

import (
	"github.com/corpetty/avalanchego/vms/secp256k1fx"
)

// MintOutput ...
type MintOutput struct {
	secp256k1fx.OutputOwners `serialize:"true"`
}
