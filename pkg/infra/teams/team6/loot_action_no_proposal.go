package team6

import (
	"infra/game/agent"
	"infra/game/commons"

	"github.com/benbjohnson/immutable"
)

// TODO
func (t6 *Team6Agent) LootActionNoProposal(baseAgent agent.BaseAgent) immutable.SortedMap[commons.ItemID, struct{}] {
	return *immutable.NewSortedMap[commons.ItemID, struct{}](nil)
}
