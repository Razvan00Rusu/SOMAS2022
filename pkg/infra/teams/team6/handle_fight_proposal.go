package team6

import (
	"infra/game/agent"
	"infra/game/decision"
	"infra/game/message"
	"math/rand"
)

func (r *Team6Agent) HandleFightProposal(
	_ message.Proposal[decision.FightAction],
	_ agent.BaseAgent,
) decision.Intent {
	intent := rand.Intn(2)
	if intent == 0 {
		return decision.Positive
	} else {
		return decision.Negative
	}
}
