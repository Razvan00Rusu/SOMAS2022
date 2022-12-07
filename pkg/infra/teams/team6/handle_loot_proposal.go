package team6

import (
	"infra/game/agent"
	"infra/game/decision"
	"infra/game/message"
)

// TODO
func (t6 *Team6Agent) HandleLootProposal(_ message.Proposal[decision.LootAction], _ agent.BaseAgent) decision.Intent {
	return decision.Abstain
}
