package minerRewardStrategies

import (
	"github.com/ethereum/go-ethereum/common"
	tmspEthTypes "github.com/tendermint/ethermint/types"
)

type RewardConstant struct {
	Etherbase      common.Address
}

func (rewardStrategy *RewardConstant) Receiver(app tmspEthTypes.TMSPEthereumApplicationInterface) common.Address {
	return rewardStrategy.Etherbase
}
