package team6

import (
	"infra/game/agent"
	"infra/game/decision"
	"infra/game/message"
)

func (r *Perry) HandleLootProposalRequest(
	msg message.Proposal[decision.LootAction],
	base agent.BaseAgent,
) bool {
	// What does this do
	return false
}
