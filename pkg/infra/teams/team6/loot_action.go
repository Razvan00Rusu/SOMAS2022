package team6

import (
	"infra/game/agent"
	"infra/game/decision"
	"infra/game/message"

	"github.com/benbjohnson/immutable"
)

func (r *Perry) LootAction(
	base agent.BaseAgent,
	proposedLoot immutable.SortedMap[string, struct{}],
	acceptedProposal message.Proposal[decision.LootAction],
) immutable.SortedMap[string, struct{}] {
	// TODO?
	return proposedLoot
}
