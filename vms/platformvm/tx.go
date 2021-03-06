// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"fmt"

	"github.com/corpetty/avalanchego/codec"
	"github.com/corpetty/avalanchego/database"
	"github.com/corpetty/avalanchego/database/versiondb"
	"github.com/corpetty/avalanchego/ids"
	"github.com/corpetty/avalanchego/snow"
	"github.com/corpetty/avalanchego/utils/crypto"
	"github.com/corpetty/avalanchego/utils/hashing"
	"github.com/corpetty/avalanchego/vms/components/verify"
	"github.com/corpetty/avalanchego/vms/secp256k1fx"
)

// UnsignedTx is an unsigned transaction
type UnsignedTx interface {
	Initialize(unsignedBytes, signedBytes []byte)
	ID() ids.ID
	UnsignedBytes() []byte
	Bytes() []byte
}

// UnsignedDecisionTx is an unsigned operation that can be immediately decided
type UnsignedDecisionTx interface {
	UnsignedTx

	// Attempts to verify this transaction with the provided state.
	SemanticVerify(vm *VM, db database.Database, stx *Tx) (
		onAcceptFunc func() error,
		err TxError,
	)
}

// UnsignedProposalTx is an unsigned operation that can be proposed
type UnsignedProposalTx interface {
	UnsignedTx

	// Attempts to verify this transaction with the provided state.
	SemanticVerify(vm *VM, db database.Database, stx *Tx) (
		onCommitDB *versiondb.Database,
		onAbortDB *versiondb.Database,
		onCommitFunc func() error,
		onAbortFunc func() error,
		err TxError,
	)
	InitiallyPrefersCommit(vm *VM) bool
}

// UnsignedAtomicTx is an unsigned operation that can be atomically accepted
type UnsignedAtomicTx interface {
	UnsignedTx

	// UTXOs this tx consumes
	InputUTXOs() ids.Set
	// Attempts to verify this transaction with the provided state.
	SemanticVerify(vm *VM, db database.Database, stx *Tx) TxError

	// Accept this transaction with the additionally provided state transitions.
	Accept(ctx *snow.Context, batch database.Batch) error
}

// Tx is a signed transaction
type Tx struct {
	// The body of this transaction
	UnsignedTx `serialize:"true" json:"unsignedTx"`

	// The credentials of this transaction
	Creds []verify.Verifiable `serialize:"true" json:"credentials"`
}

// Sign this transaction with the provided signers
func (tx *Tx) Sign(c codec.Manager, signers [][]*crypto.PrivateKeySECP256K1R) error {
	unsignedBytes, err := c.Marshal(codecVersion, &tx.UnsignedTx)
	if err != nil {
		return fmt.Errorf("couldn't marshal UnsignedTx: %w", err)
	}

	// Attach credentials
	hash := hashing.ComputeHash256(unsignedBytes)
	for _, keys := range signers {
		cred := &secp256k1fx.Credential{
			Sigs: make([][crypto.SECP256K1RSigLen]byte, len(keys)),
		}
		for i, key := range keys {
			sig, err := key.SignHash(hash) // Sign hash
			if err != nil {
				return fmt.Errorf("problem generating credential: %w", err)
			}
			copy(cred.Sigs[i][:], sig)
		}
		tx.Creds = append(tx.Creds, cred) // Attach credential
	}

	signedBytes, err := c.Marshal(codecVersion, tx)
	if err != nil {
		return fmt.Errorf("couldn't marshal ProposalTx: %w", err)
	}
	tx.Initialize(unsignedBytes, signedBytes)
	return nil
}
