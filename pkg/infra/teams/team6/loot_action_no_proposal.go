package team6

import (
	"infra/game/agent"
	"infra/game/commons"

	"github.com/benbjohnson/immutable"
)

func (r *Perry) LootActionNoProposal(
	base agent.BaseAgent,
) immutable.SortedMap[commons.ItemID, struct{}] {
	return *immutable.NewSortedMap[commons.ItemID, struct{}](nil)
}
