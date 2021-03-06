// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package node

import (
	"time"

	"github.com/corpetty/avalanchego/database"
	"github.com/corpetty/avalanchego/genesis"
	"github.com/corpetty/avalanchego/ids"
	"github.com/corpetty/avalanchego/nat"
	"github.com/corpetty/avalanchego/snow/consensus/avalanche"
	"github.com/corpetty/avalanchego/snow/networking/benchlist"
	"github.com/corpetty/avalanchego/snow/networking/router"
	"github.com/corpetty/avalanchego/utils"
	"github.com/corpetty/avalanchego/utils/dynamicip"
	"github.com/corpetty/avalanchego/utils/logging"
	"github.com/corpetty/avalanchego/utils/timer"
)

// Config contains all of the configurations of an Avalanche node.
type Config struct {
	genesis.Params

	// Genesis information
	GenesisBytes []byte
	AvaxAssetID  ids.ID

	// protocol to use for opening the network interface
	Nat nat.Router

	// Attempted NAT Traversal did we attempt
	AttemptedNATTraversal bool

	// ID of the network this node should connect to
	NetworkID uint32

	// Assertions configuration
	EnableAssertions bool

	// Crypto configuration
	EnableCrypto bool

	// Database to use for the node
	DB database.Database

	// Staking configuration
	StakingIP             utils.DynamicIPDesc
	EnableP2PTLS          bool
	EnableStaking         bool
	StakingKeyFile        string
	StakingCertFile       string
	DisabledStakingWeight uint64

	// Throttling
	MaxNonStakerPendingMsgs uint32
	StakerMSGPortion        float64
	StakerCPUPortion        float64
	SendQueueSize           uint32
	MaxPendingMsgs          uint32

	// Network configuration
	NetworkConfig timer.AdaptiveTimeoutConfig

	// Benchlist Configuration
	BenchlistConfig benchlist.Config

	// Bootstrapping configuration
	BootstrapPeers []*Peer

	// HTTP configuration
	HTTPHost string
	HTTPPort uint16

	HTTPSEnabled        bool
	HTTPSKeyFile        string
	HTTPSCertFile       string
	APIRequireAuthToken bool
	APIAuthPassword     string

	// Enable/Disable APIs
	AdminAPIEnabled    bool
	InfoAPIEnabled     bool
	KeystoreAPIEnabled bool
	MetricsAPIEnabled  bool
	HealthAPIEnabled   bool

	// Logging configuration
	LoggingConfig logging.Config

	// Plugin directory
	PluginDir string

	// Consensus configuration
	ConsensusParams avalanche.Parameters

	// Throughput configuration
	ThroughputPort          uint16
	ThroughputServerEnabled bool

	// IPC configuration
	IPCAPIEnabled      bool
	IPCPath            string
	IPCDefaultChainIDs []string

	// Router that is used to handle incoming consensus messages
	ConsensusRouter          router.Router
	ConsensusGossipFrequency time.Duration
	ConsensusShutdownTimeout time.Duration

	// Dynamic Update duration for IP or NAT traversal
	DynamicUpdateDuration time.Duration

	DynamicPublicIPResolver dynamicip.Resolver

	// Throttling incoming connections
	ConnMeterResetDuration time.Duration
	ConnMeterMaxConns      int

	// Subnet Whitelist
	WhitelistedSubnets ids.Set

	// Restart on disconnect settings
	RestartOnDisconnected      bool
	DisconnectedCheckFreq      time.Duration
	DisconnectedRestartTimeout time.Duration

	// Coreth
	CorethConfig string
}
