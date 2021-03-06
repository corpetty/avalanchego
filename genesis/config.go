// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/corpetty/avalanchego/ids"
	"github.com/corpetty/avalanchego/utils/constants"
	"github.com/corpetty/avalanchego/utils/formatting"
	"github.com/corpetty/avalanchego/utils/wrappers"

	safemath "github.com/corpetty/avalanchego/utils/math"
)

// LockedAmount ...
type LockedAmount struct {
	Amount   uint64 `json:"amount"`
	Locktime uint64 `json:"locktime"`
}

// Allocation ...
type Allocation struct {
	ETHAddr        ids.ShortID    `json:"ethAddr"`
	AVAXAddr       ids.ShortID    `json:"avaxAddr"`
	InitialAmount  uint64         `json:"initialAmount"`
	UnlockSchedule []LockedAmount `json:"unlockSchedule"`
}

// Unparse ...
func (a Allocation) Unparse(networkID uint32) (UnparsedAllocation, error) {
	ua := UnparsedAllocation{
		InitialAmount:  a.InitialAmount,
		UnlockSchedule: a.UnlockSchedule,
		ETHAddr:        "0x" + hex.EncodeToString(a.ETHAddr.Bytes()),
	}
	avaxAddr, err := formatting.FormatAddress(
		"X",
		constants.GetHRP(networkID),
		a.AVAXAddr.Bytes(),
	)
	ua.AVAXAddr = avaxAddr
	return ua, err
}

// Staker ...
type Staker struct {
	NodeID        ids.ShortID `json:"nodeID"`
	RewardAddress ids.ShortID `json:"rewardAddress"`
	DelegationFee uint32      `json:"delegationFee"`
}

// Unparse ...
func (s Staker) Unparse(networkID uint32) (UnparsedStaker, error) {
	avaxAddr, err := formatting.FormatAddress(
		"X",
		constants.GetHRP(networkID),
		s.RewardAddress.Bytes(),
	)
	return UnparsedStaker{
		NodeID:        s.NodeID.PrefixedString(constants.NodeIDPrefix),
		RewardAddress: avaxAddr,
		DelegationFee: s.DelegationFee,
	}, err
}

// Config contains the genesis addresses used to construct a genesis
type Config struct {
	NetworkID uint32 `json:"networkID"`

	Allocations []Allocation `json:"allocations"`

	StartTime                  uint64        `json:"startTime"`
	InitialStakeDuration       uint64        `json:"initialStakeDuration"`
	InitialStakeDurationOffset uint64        `json:"initialStakeDurationOffset"`
	InitialStakedFunds         []ids.ShortID `json:"initialStakedFunds"`
	InitialStakers             []Staker      `json:"initialStakers"`

	CChainGenesis string `json:"cChainGenesis"`

	Message string `json:"message"`
}

// Unparse ...
func (c Config) Unparse() (UnparsedConfig, error) {
	uc := UnparsedConfig{
		NetworkID:                  c.NetworkID,
		Allocations:                make([]UnparsedAllocation, len(c.Allocations)),
		StartTime:                  c.StartTime,
		InitialStakeDuration:       c.InitialStakeDuration,
		InitialStakeDurationOffset: c.InitialStakeDurationOffset,
		InitialStakedFunds:         make([]string, len(c.InitialStakedFunds)),
		InitialStakers:             make([]UnparsedStaker, len(c.InitialStakers)),
		CChainGenesis:              c.CChainGenesis,
		Message:                    c.Message,
	}
	for i, a := range c.Allocations {
		ua, err := a.Unparse(uc.NetworkID)
		if err != nil {
			return uc, err
		}
		uc.Allocations[i] = ua
	}
	for i, isa := range c.InitialStakedFunds {
		avaxAddr, err := formatting.FormatAddress(
			"X",
			constants.GetHRP(uc.NetworkID),
			isa.Bytes(),
		)
		if err != nil {
			return uc, err
		}
		uc.InitialStakedFunds[i] = avaxAddr
		fmt.Println(avaxAddr)
	}
	for i, is := range c.InitialStakers {
		uis, err := is.Unparse(c.NetworkID)
		if err != nil {
			return uc, err
		}
		uc.InitialStakers[i] = uis
	}

	return uc, nil
}

// InitialSupply ...
func (c *Config) InitialSupply() (uint64, error) {
	initialSupply := uint64(0)
	for _, allocation := range c.Allocations {
		newInitialSupply, err := safemath.Add64(initialSupply, allocation.InitialAmount)
		if err != nil {
			return 0, err
		}
		for _, unlock := range allocation.UnlockSchedule {
			newInitialSupply, err = safemath.Add64(newInitialSupply, unlock.Amount)
			if err != nil {
				return 0, err
			}
		}
		initialSupply = newInitialSupply
	}
	return initialSupply, nil
}

var (
	// MainnetConfig is the config that should be used to generate the mainnet
	// genesis.
	MainnetConfig Config

	// FujiConfig is the config that should be used to generate the fuji
	// genesis.
	FujiConfig Config

	// LocalConfig is the config that should be used to generate a local
	// genesis.
	LocalConfig Config

	// StatalancheConfig is the config that should be used to generate the Statalanche
	// genesis
	StatalancheConfig Config
)

func init() {
	unparsedMainnetConfig := UnparsedConfig{}
	unparsedFujiConfig := UnparsedConfig{}
	unparsedLocalConfig := UnparsedConfig{}
	unparsedStatalancheConfig := UnparsedConfig{}

	errs := wrappers.Errs{}
	errs.Add(
		json.Unmarshal([]byte(mainnetGenesisConfigJSON), &unparsedMainnetConfig),
		json.Unmarshal([]byte(fujiGenesisConfigJSON), &unparsedFujiConfig),
		json.Unmarshal([]byte(localGenesisConfigJSON), &unparsedLocalConfig),
		json.Unmarshal([]byte(statalancheGenesisConfigJSON), &unparsedStatalancheConfig),
	)
	if errs.Errored() {
		panic(errs.Err)
	}

	mainnetConfig, err := unparsedMainnetConfig.Parse()
	errs.Add(err)
	MainnetConfig = mainnetConfig

	fujiConfig, err := unparsedFujiConfig.Parse()
	errs.Add(err)
	FujiConfig = fujiConfig

	localConfig, err := unparsedLocalConfig.Parse()
	errs.Add(err)
	LocalConfig = localConfig

	statalancheConfig, err := unparsedStatalancheConfig.Parse()
	errs.Add(err)
	StatalancheConfig = statalancheConfig

	if errs.Errored() {
		panic(errs.Err)
	}
}

// GetConfig ...
func GetConfig(networkID uint32) *Config {
	switch networkID {
	case constants.MainnetID:
		return &MainnetConfig
	case constants.FujiID:
		return &FujiConfig
	case constants.LocalID:
		return &LocalConfig
	case constants.StatalancheID:
		return &StatalancheConfig
	default:
		tempConfig := LocalConfig
		tempConfig.NetworkID = networkID
		return &tempConfig
	}
}

// GetConfigFile loads a *Config from a provided
// filepath.
func GetConfigFile(filepath string) (*Config, error) {
	b, err := ioutil.ReadFile(path.Clean(filepath))
	if err != nil {
		return nil, fmt.Errorf("unable to load file %s: %w", filepath, err)
	}

	var unparsedConfig UnparsedConfig
	if err := json.Unmarshal(b, &unparsedConfig); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	config, err := unparsedConfig.Parse()
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	return &config, nil
}
