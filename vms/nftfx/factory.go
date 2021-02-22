package nftfx

import (
	"github.com/corpetty/avalanchego/ids"
	"github.com/corpetty/avalanchego/snow"
)

// ID that this Fx uses when labeled
var (
	ID = ids.ID{'n', 'f', 't', 'f', 'x'}
)

// Factory ...
type Factory struct{}

// New ...
func (f *Factory) New(*snow.Context) (interface{}, error) { return &Fx{}, nil }
