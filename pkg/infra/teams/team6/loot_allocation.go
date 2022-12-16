package team6

import (
	"infra/game/agent"
	"infra/game/decision"
	"infra/game/message"

	"github.com/benbjohnson/immutable"
)

func (r *Perry) LootAllocation(
	base agent.BaseAgent,
	proposal message.Proposal[decision.LootAction],
	proposedAllocations immutable.Map[string, immutable.SortedMap[string, struct{}]],
) immutable.Map[string, immutable.SortedMap[string, struct{}]] {
	panic("Not Implemented")
}
