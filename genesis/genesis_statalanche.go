// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/ava-labs/avalanchego/utils/units"
)

var (
	statalancheGenesisConfigJSON = `{
		"networkID": 115110116,
		"allocations": [
			{
				"avaxAddr": "X-avax1ue6kq32j382vxp2gq8exzztyn2mg6htg9l4kgk",
				"ethAddr": "0xe027688a57c4A6Fb2708343cF330aaeB8fe594bb",
				"initialAmount": 10000000000000000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000
					}
				]
			},
			{
				"ethAddr": "0xe027688a57c4A6Fb2708343cF330aaeB8fe594bb",
				"avaxAddr": "X-avax1ng093n7uceg5g8lmp450kwy64j9newc5m0yjun",
				"initialAmount": 300000000000000000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000
					}
				]
			},
			{
				"ethAddr": "0xe027688a57c4A6Fb2708343cF330aaeB8fe594bb",
				"avaxAddr": "X-avax1xfwyt3hwn8xllgm5uw03cymk9sxuxpdpw3y48e",
				"initialAmount": 10000000000000000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000
					}
				]
			},
			{
				"ethAddr": "0xe027688a57c4A6Fb2708343cF330aaeB8fe594bb",
				"avaxAddr": "X-avax1texzpd6ea9x4al2aavrdkyjlwt99pjfsfjtxal",
				"initialAmount": 10000000000000000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000
					}
				]
			},
			{
				"ethAddr": "0xe027688a57c4A6Fb2708343cF330aaeB8fe594bb",
				"avaxAddr": "X-avax1sd0tw9d3wx5xu4vk7r2qqtwc7lyzkw35mwzmvm",
				"initialAmount": 10000000000000000,
				"unlockSchedule": [
					{
						"amount": 10000000000000000
					}
				]
			}
		],
		"startTime": 1599696000,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [
			"X-avax1ue6kq32j382vxp2gq8exzztyn2mg6htg9l4kgk",
			"X-avax1ng093n7uceg5g8lmp450kwy64j9newc5m0yjun",
			"X-avax1xfwyt3hwn8xllgm5uw03cymk9sxuxpdpw3y48e",
			"X-avax1texzpd6ea9x4al2aavrdkyjlwt99pjfsfjtxal",
			"X-avax1sd0tw9d3wx5xu4vk7r2qqtwc7lyzkw35mwzmvm"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-E349V27puxyemQTc6QDDGaLD1kEF1GJfa",
				"rewardAddress": "X-avax1ue6kq32j382vxp2gq8exzztyn2mg6htg9l4kgk",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-Pvfe5vYMUMecgXWcbtmSjqVxcisQJuHty",
				"rewardAddress": "X-avax1ng093n7uceg5g8lmp450kwy64j9newc5m0yjun",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-PCe2PQb5wwXT7JUsFQjrqb8bVhRinnKoW",
				"rewardAddress": "X-avax1xfwyt3hwn8xllgm5uw03cymk9sxuxpdpw3y48e",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-MAgi5KLrdqjoPFMHEySNcFS8Ax7EsLB1f",
				"rewardAddress": "X-avax1texzpd6ea9x4al2aavrdkyjlwt99pjfsfjtxal",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-DjkQepxbpuVM9yZrYmekhKkKwSeQa3MR",
				"rewardAddress": "X-avax1sd0tw9d3wx5xu4vk7r2qqtwc7lyzkw35mwzmvm",
				"delegationFee": 200000
			}
		],
		"cChainGenesis": "{\"config\":{\"chainId\":13375,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"0100000000000000000000000000000000000000\":{\"code\":\"0x7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80631e010439146042578063b6510bb314606e575b600080fd5b605c60048036036020811015605657600080fd5b503560b1565b60408051918252519081900360200190f35b818015607957600080fd5b5060af60048036036080811015608e57600080fd5b506001600160a01b03813516906020810135906040810135906060013560b6565b005b30cd90565b836001600160a01b031681836108fc8690811502906040516000604051808303818888878c8acf9550505050505015801560f4573d6000803e3d6000fd5b505050505056fea26469706673582212201eebce970fe3f5cb96bf8ac6ba5f5c133fc2908ae3dcd51082cfee8f583429d064736f6c634300060a0033\",\"balance\":\"0x0\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "It is all about the State of Us"
	}`

	// StatalancheParams are the params used for local networks
	StatalancheParams = Params{
		TxFee:                units.MilliAvax,
		CreationTxFee:        10 * units.MilliAvax,
		UptimeRequirement:    .6, // 60%
		MinValidatorStake:    1 * units.Avax,
		MaxValidatorStake:    3 * units.MegaAvax,
		MinDelegatorStake:    1 * units.Avax,
		MinDelegationFee:     20000, // 2%
		MinStakeDuration:     24 * time.Hour,
		MaxStakeDuration:     365 * 24 * time.Hour,
		StakeMintingPeriod:   365 * 24 * time.Hour,
		EpochFirstTransition: time.Unix(1607626800, 0),
		EpochDuration:        30 * time.Minute,
		ApricotPhase0Time:    time.Date(2020, 12, 5, 5, 00, 0, 0, time.UTC),
	}
)