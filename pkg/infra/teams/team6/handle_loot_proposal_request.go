package team6

import (
	"infra/game/agent"
	"infra/game/decision"
	"infra/game/message"
)

// TODO
func (t6 *Team6Agent) HandleLootProposalRequest(_ message.Proposal[decision.LootAction], _ agent.BaseAgent) bool {
	return true
}
