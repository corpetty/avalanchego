// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/corpetty/avalanchego/snow/consensus/snowball"
	"github.com/corpetty/avalanchego/snow/consensus/snowman"
	"github.com/corpetty/avalanchego/snow/engine/snowman/bootstrap"
)

// Config wraps all the parameters needed for a snowman engine
type Config struct {
	bootstrap.Config

	Params    snowball.Parameters
	Consensus snowman.Consensus
}
