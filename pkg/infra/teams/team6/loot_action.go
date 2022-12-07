package team6

import (
	"infra/game/agent"
	"infra/game/commons"

	"github.com/benbjohnson/immutable"
)

// TODO
func (t6 *Team6Agent) LootAction(baseAgent agent.BaseAgent, proposedLoot immutable.SortedMap[commons.ItemID, struct{}]) immutable.SortedMap[commons.ItemID, struct{}] {
	return proposedLoot
}
